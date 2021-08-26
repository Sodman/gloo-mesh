package docsgen

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"

	"github.com/Masterminds/semver"
	"github.com/google/go-github/v32/github"
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/stoewer/go-strcase"
)

var (
	helmDocsDir = "content/reference/helm"

	ossFileMapping = map[string]string{
		"codegen/helm/gloo_mesh_helm_values_reference.md":  "%s/%s/gloo_mesh.md",
		"codegen/helm/cert_agent_helm_values_reference.md": "%s/%s/cert_agent.md",
	}

	enterpriseFileMapping = map[string]string{
		"enterprise-networking/codegen/helm/enterprise_networking_helm_values_reference.md": "%s/%s/enterprise_networking.md",
		"enterprise-networking/codegen/helm/enterprise_agent_helm_values_reference.md":      "%s/%s/enterprise_agent.md",
		"rbac-webhook/codegen/chart/rbac_webhook_helm_values_reference.md":                  "%s/%s/rbac_webhook.md",
		"gloo-mesh-ui/codegen/helm/dashboard_helm_values_reference.md":                      "%s/%s/dashboard_helm_values_reference.md",
		"gloo-mesh-ui/codegen/helm/redis_helm_values_reference.md":                          "%s/%s/redis_helm_values_reference.md",
	}

	helmValuesIndex = `
---
title: "%s"
description: Reference for Helm values. 
weight: 2
---

The following pages provide Helm value reference documentation for various Gloo Mesh components, including:

1. **Open source Gloo Mesh**: the OSS version of Gloo Mesh
2. **Enterprise Networking (enterprise only)**: The management plane of Gloo Mesh Enterprise, deployed on the management cluster.
3. **Enterprise Agent (enterprise only)**: The agent of Gloo Mesh Enterprise, deployed on each remote cluster.
4. **RBAC Webhook (enterprise only)**: The Kubernetes webhook that enforces Gloo Mesh Enterprise's role-based API.
5. **Gloo Mesh UI (enterprise only)**: The UI for Gloo Mesh Enterprise, which includes the Dashboard and Redis subcharts.

Note that when you provide Helm values for the [bundled Gloo Mesh Enterprise chart](https://storage.googleapis.com/gloo-mesh-enterprise/gloo-mesh-enterprise),
values for each subchart must be prefixed accordingly:

1. Values for the RBAC Webhook must be prefixed with "rbac-webhook".
2. Values for Enterprise Networking must be prefixed with "enterprise-networking".
3. Values for the Gloo Mesh UI must be prefixed with "gloo-mesh-ui".
  - Values for the Dashboard (a subchart of the Gloo Mesh UI chart) must be prefixed with "dashboard"
  - Values for the Redis Dashboard (a subchart of the Gloo Mesh UI chart) must be prefixed with "redis-dashboard"


The following is an example of how to set values for subcharts:


> helm install gloo-mesh-enterprise gloo-mesh-enterprise/gloo-mesh-enterprise \
> --namespace gloo-mesh \
> --set licenseKey=${GLOO_MESH_LICENSE_KEY}  \
> --set rbac-webhook.enabled=true
> --set enterprise-networking.enterpriseNetworking.floatingUserId=true \
> --set gloo-mesh-ui.dashboard.floatingUserId.floatingUserId=true \
> --set gloo-mesh-ui.redis-dashboard.redisDashboard.floatingUserId=true

{{%% children description="true" %%}}
`
	numberMatcher = regexp.MustCompile("[0-9]+")
)

func copyHelmValuesDocsForAllCharts(client *github.Client, rootDir string) error {
	// flush root directory for idempotence
	helmDocsDir := filepath.Join(rootDir, helmDocsDir)
	os.RemoveAll(helmDocsDir)
	os.MkdirAll(helmDocsDir, 0755)

	// create root index
	if err := createFileIfNotExists(helmDocsDir+"/"+"_index.md", fmt.Sprintf(helmValuesIndex, "Helm Values Reference")); err != nil {
		return eris.Errorf("error creating Helm values index file: %v", err)
	}

	// Gloo Mesh OSS
	if err := copyHelmValuesDocsForComponent(
		client,
		rootDir,
		"Gloo Mesh",
		GlooMeshRepoName,
		"v1.0.0",
		ossFileMapping,
	); err != nil {
		return err
	}

	// Gloo Mesh Enterprise
	if err := copyHelmValuesDocsForComponent(
		client,
		rootDir,
		"Gloo Mesh Enterprise",
		GlooMeshEnterpriseRepoName,
		"v1.0.0",
		enterpriseFileMapping,
	); err != nil {
		return err
	}

	return nil
}

// fetch Helm Values documentation from repo up to and including the version specified by earliestVerison
// fileMapping specifies a mapping from the file path in the origin repo to the file path in this repo
func copyHelmValuesDocsForComponent(
	client *github.Client,
	rootDir string,
	componentName string,
	repoName string,
	earliestVersion string,
	fileMapping map[string]string,
) error {
	// flush directory for idempotence
	helmDocsDir := filepath.Join(rootDir, helmDocsDir, strcase.SnakeCase(componentName))
	os.RemoveAll(helmDocsDir)
	os.MkdirAll(helmDocsDir, 0755)

	if err := createFileIfNotExists(helmDocsDir+"/"+"_index.md", fmt.Sprintf(helmValuesIndex, componentName)); err != nil {
		return eris.Errorf("error creating Helm values index file: %v", err)
	}

	// include Helm values docs for all versions > earliestVersion
	releases, _, err := client.Repositories.ListReleases(
		context.Background(),
		GithubOrg,
		repoName,
		&github.ListOptions{Page: 0, PerPage: 1000000},
	)
	if err != nil {
		return eris.Errorf("error listing releases: %v", err)
	}

	// the github API returns releases sorted by release date, so we need to sort by version in order to enforce the earliest version lower bound
	var versions []*semver.Version
	for _, release := range releases {
		tagName := release.GetTagName()
		version, err := semver.NewVersion(tagName)
		if err != nil {
			return err
		}
		var modifiedVersion semver.Version
		if version.Prerelease() != "" {
			// semver's comparison function will not put 'beta9' ahead of 'beta10', so we modify the
			// prerelease text to just the number in the prerelease tag.
			match := numberMatcher.FindAllString(version.Prerelease(), -1)
			modifiedVersion, err = version.SetPrerelease(match[0])
			if err != nil {
				return err
			}
			versions = append(versions, &modifiedVersion)
		} else {
			versions = append(versions, version)
		}
	}
	sort.Sort(sort.Reverse(semver.Collection(versions)))
	latestPerMinorVersions := getLatestPerMinorVersion(versions)

	earliestVersionSemver, err := semver.NewVersion(earliestVersion)
	if err != nil {
		return err
	}
	latestVersionSemver, err := semver.NewVersion(earliestVersion)
	if err != nil {
		return err
	}

	tags := make(map[string]string, 0)
	for _, version := range latestPerMinorVersions {
		tags[version.Original()] = fmt.Sprintf("%d.%d", version.Major(), version.Minor())
		if version.GreaterThan(latestVersionSemver) {
			latestVersionSemver = version
		}
		if version.LessThan(earliestVersionSemver) || version.Equal(earliestVersionSemver) {
			break
		}
	}
	tags[latestVersionSemver.Original()] = "latest"

	for tag, tagPath := range tags {
		contextutils.LoggerFrom(context.Background()).Infof("copying Helm values docs from %s/%s for release %s", GithubOrg, repoName, tag)

		if err := os.Mkdir(helmDocsDir+"/"+tagPath, os.ModePerm); err != nil {
			return eris.Errorf("error creating Helm docs directories: %v", err)
		}

		if err := createFileIfNotExists(helmDocsDir+"/"+tagPath+"/"+"_index.md", fmt.Sprintf(helmValuesIndex, tag)); err != nil {
			return eris.Errorf("error creating Helm values index file: %v", err)
		}

		for src, dest := range fileMapping {
			dest = fmt.Sprintf(dest, helmDocsDir, tagPath)
			if err := copyHelmValuesDocs(client, GithubOrg, repoName, tag, src, dest); err != nil {
				return err
			}
		}
	}

	return nil
}

// returns the latest patch version for each minor version
// expects versions to be sorted in reverse order
func getLatestPerMinorVersion(sortedVersions []*semver.Version) []*semver.Version {
	var latestVersions []*semver.Version

	latestVersionForMinor, _ := semver.NewVersion("1.999999999.0")
	for _, version := range sortedVersions {
		if version.Minor() < latestVersionForMinor.Minor() {
			if version.Prerelease() != "" {
				// semver's comparison function will not put 'beta9' ahead of 'beta10', so we revert the modified
				// prerelease text to the original beta-number pattern.
				origVersion, _ := version.SetPrerelease("beta" + version.Prerelease())
				latestVersions = append(latestVersions, &origVersion)
				latestVersionForMinor = &origVersion
			} else {
				latestVersions = append(latestVersions, version)
				latestVersionForMinor = version
			}
		}
	}

	return latestVersions
}

func copyHelmValuesDocs(client *github.Client, org, repo, tag, path, destinationFile string) error {
	contents, _, resp, err := client.Repositories.GetContents(context.Background(), org, repo, path, &github.RepositoryContentGetOptions{
		Ref: tag,
	})

	// return error if expected doc files aren't found
	if err != nil && resp != nil && resp.StatusCode == 404 {
		// 404 is expected for Helm values docs introduced in later versions of Gloo Mesh
		logrus.Warnf("missing Helm values file \"%s\" for version \"%s\"", path, tag)
		return nil
	} else if err != nil {
		return eris.Errorf("error fetching Helm values doc: %v", err)
	}

	decodedContents, err := base64.StdEncoding.DecodeString(*contents.Content)
	if err != nil {
		return eris.Errorf("error fetching Helm values doc: %v", err)
	}

	return createFileIfNotExists(destinationFile, string(decodedContents))
}

// create file with contents, create the file if it doesn't exist
func createFileIfNotExists(fname, contents string) error {
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		return eris.Errorf("error copying Helm values reference doc: %v", err)
	}

	if _, err = f.Write([]byte(contents)); err != nil {
		return err
	}

	return nil
}
