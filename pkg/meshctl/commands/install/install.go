package install

import (
	"context"
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"
	"github.com/solo-io/gloo-mesh/pkg/common/defaults"
	cliversion "github.com/solo-io/gloo-mesh/pkg/common/version"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/enterprise"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/install/gloomesh"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/install/helm"
	installutils "github.com/solo-io/gloo-mesh/pkg/meshctl/install/utils"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/registration"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/utils"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/validation"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/validation/checks"
	validationconsts "github.com/solo-io/gloo-mesh/pkg/meshctl/validation/consts"
	"github.com/solo-io/skv2/pkg/crdutils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Command(ctx context.Context, globalFlags *utils.GlobalFlags) *cobra.Command {
	opts := &Options{GlobalFlags: globalFlags}
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install Gloo Mesh",
		Long: `Install the Gloo Mesh management plan to a Kubernetes cluster.

Go to https://www.solo.io/products/gloo-mesh/ to learn more about the
difference between the editions.
`,
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			// Update the meshctl config file
			utils.UpdateMeshctlConfigWithInstallInfo(opts.KubeCfgPath, opts.KubeContext)
		},
	}

	cmd.AddCommand(
		communityCommand(ctx, opts),
		enterpriseCommand(ctx, opts),
	)

	opts.addToFlags(cmd.PersistentFlags())

	return cmd
}

type Options struct {
	*utils.GlobalFlags
	DryRun bool

	KubeCfgPath     string
	KubeContext     string
	Namespace       string
	ChartPath       string
	ChartValuesFile string
	ExtraHelmValues []string
	Version         string
	ReleaseName     string
	AgentChartPath  string
	AgentValuesPath string

	Register      bool
	ClusterName   string
	ClusterDomain string
}

func (o *Options) addToFlags(flags *pflag.FlagSet) {
	utils.AddManagementKubeconfigFlags(&o.KubeCfgPath, &o.KubeContext, flags)
	flags.BoolVarP(&o.DryRun, "dry-run", "d", false, "Output installation manifest")
	flags.StringVar(&o.Namespace, "namespace", defaults.DefaultPodNamespace, "Namespace in which to install Gloo Mesh")
	flags.StringVar(&o.ChartPath, "chart-file", "", "Path to a local Helm chart for installing Gloo Mesh.\nIf unset, this command will install Gloo Mesh from the publicly released Helm chart.")
	flags.StringVar(&o.ChartValuesFile, "chart-values-file", "", "File containing value overrides for the Gloo Mesh Helm chart")
	flags.StringVar(&o.Version, "version", "", "Version to install.\nCommunity defaults to meshctl version, enterprise defaults to latest stable")
	flags.StringArrayVar(&o.ExtraHelmValues, "set", []string{}, "Extra helm values for the Gloo Mesh chart.")
	flags.BoolVarP(&o.Register, "register", "r", false, "Also register the cluster")
	flags.StringVar(&o.ClusterName, "cluster-name", "mgmt-cluster",
		"Name with which to register the cluster running Gloo Mesh, only applies if --register is also set")
	flags.StringVar(&o.ClusterDomain, "cluster-domain", defaults.DefaultClusterDomain, "The cluster domain used by the Kubernetes DNS Service in the registered cluster. \nRead more: https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/")
}

func (o *Options) getInstaller(chartUriTemplate string) helm.Installer {
	// User-specified ChartPath takes precedence over specified Version.
	if o.ChartPath == "" {
		o.ChartPath = fmt.Sprintf(chartUriTemplate, strings.TrimPrefix(o.Version, "v"))
	}

	logrus.Debugf("installing chart from %s", o.ChartPath)
	return helm.Installer{
		ChartUri:    o.ChartPath,
		ValuesFile:  o.ChartValuesFile,
		KubeConfig:  o.KubeCfgPath,
		KubeContext: o.KubeContext,
		Namespace:   o.Namespace,
		ReleaseName: o.ReleaseName,
		Values:      getStringMap(o.ExtraHelmValues),
		Verbose:     o.Verbose,
		DryRun:      o.DryRun,
	}
}

func (o *Options) getRegistrationOptions() registration.Options {
	return registration.Options{
		KubeConfigPath:         o.KubeCfgPath,
		MgmtContext:            o.KubeContext,
		RemoteContext:          o.KubeContext,
		ClusterName:            o.ClusterName,
		MgmtNamespace:          o.Namespace,
		RemoteNamespace:        o.Namespace,
		AgentChartPathOverride: o.AgentChartPath,
		AgentChartValuesPath:   o.AgentValuesPath,
		ClusterDomain:          o.ClusterDomain,
		Version:                o.Version,
		Verbose:                o.Verbose,
	}
}

func communityCommand(ctx context.Context, installOpts *Options) *cobra.Command {
	opts := CommunityOptions{Options: installOpts}
	cmd := &cobra.Command{
		Use:   "community",
		Short: "Install Gloo Mesh Community",
		Example: `  # Install to the currently selected Kubernetes context
  meshctl install community

  # Install to and Register the currently selected Kubernetes context
  meshctl install community --register

  # Install to a different context
  meshctl install --kubecontext=some-context community`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return InstallCommunity(ctx, opts)
		},
	}

	opts.addToFlags(cmd.Flags())
	cmd.SilenceUsage = true
	return cmd
}

type CommunityOptions struct {
	*Options
	AgentCrdsChartPath string
	ApiServerAddress   string
}

func (o *CommunityOptions) addToFlags(flags *pflag.FlagSet) {
	flags.StringVar(&o.ReleaseName, "release-name", gloomesh.GlooMeshReleaseName, "Helm release name")
	flags.StringVar(&o.ApiServerAddress, "api-server-address", "", "Swap out the address of the remote cluster's k8s API server for the value of this flag.\nSet this flag when the address of the cluster domain used by the Gloo Mesh is different than that specified in the local kubeconfig.")
	flags.StringVar(&o.AgentCrdsChartPath, "agent-crds-chart-file", "", "Path to a local Helm chart for installing CRDs needed by remote agents.\nIf unset, this command will install the agent CRDs from the publicly released Helm chart.")
	flags.StringVar(&o.AgentChartPath, "cert-agent-chart-file", "",
		"Path to a local Helm chart for installing the Certificate Agent.\n"+
			"If unset, this command will install the Certificate Agent from the publicly released Helm chart.",
	)
	flags.StringVar(&o.AgentValuesPath, "cert-agent-chart-values", "",
		"Path to a Helm values.yaml file for customizing the installation of the Certificate Agent.\n"+
			"If unset, this command will install the Certificate Agent with default Helm values.",
	)
}

func (o CommunityOptions) getInstaller() helm.Installer {
	return o.Options.getInstaller(gloomesh.GlooMeshChartUriTemplate)
}

func (o CommunityOptions) getRegistrationOptions() registration.Options {
	reg := o.Options.getRegistrationOptions()
	reg.AgentCrdsChartPath = o.AgentCrdsChartPath
	reg.ApiServerAddress = o.ApiServerAddress
	return reg
}

func InstallCommunity(ctx context.Context, opts CommunityOptions) error {
	if opts.Version == "" {
		opts.Version = cliversion.Version
	}
	logrus.Info("Installing Helm chart")
	if err := opts.getInstaller().InstallChart(ctx); err != nil {
		return eris.Wrap(err, "installing gloo-mesh")
	}

	if opts.Register && !opts.DryRun {
		logrus.Info("Registering cluster")
		registrant, err := registration.NewRegistrant(opts.getRegistrationOptions())
		if err != nil {
			return eris.Wrap(err, "initializing registrant")
		}
		if err := registrant.RegisterCluster(ctx); err != nil {
			return eris.Wrap(err, "registering management-plane cluster")
		}
	}

	return nil
}

func enterpriseCommand(ctx context.Context, installOpts *Options) *cobra.Command {
	opts := EnterpriseOptions{Options: installOpts}
	cmd := &cobra.Command{
		Use:   "enterprise",
		Short: "Install Gloo Mesh Enterprise (requires a license)",
		Example: `  # Install to the currently selected Kubernetes context
  meshctl install enterprise --license=<my_license>

  # Install to and Register the currently selected Kubernetes context
  meshctl install enterprise --license=<my_license> --Register

  # Don't install the UI
  meshctl install enterprise --license=<my_license> --skip-ui`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return InstallEnterprise(ctx, opts)
		},
	}

	opts.addToFlags(cmd.Flags())
	cmd.SilenceUsage = true
	return cmd
}

type EnterpriseOptions struct {
	*Options
	LicenseKey         string
	SkipUI             bool
	IncludeRBAC        bool
	RelayServerAddress string
	SkipChecks         bool
}

func (o *EnterpriseOptions) addToFlags(flags *pflag.FlagSet) {
	flags.StringVar(&o.ReleaseName, "release-name", gloomesh.GlooMeshReleaseName, "Helm release name")
	flags.StringVar(&o.LicenseKey, "license", "", "Gloo Mesh Enterprise license key (required)")
	flags.BoolVar(&o.SkipUI, "skip-ui", false, "Skip installation of the Gloo Mesh UI")
	flags.BoolVar(&o.IncludeRBAC, "include-rbac", false, "Install the RBAC Webhook")
	flags.StringVar(&o.AgentChartPath, "enterprise-agent-chart-file", "",
		"Path to a local Helm chart for installing the Enterprise Agent.\n"+
			"If unset, this command will install the Enterprise Agent from the publicly released Helm chart.",
	)
	flags.StringVar(&o.AgentValuesPath, "enterprise-agent-chart-values", "",
		"Path to a Helm values.yaml file for customizing the installation of the Enterprise Agent.\n"+
			"If unset, this command will install the Enterprise Agent with default Helm values.",
	)
	flags.StringVar(&o.RelayServerAddress, "relay-server-address", "", "The address that the enterprise agent will communicate with the relay server via. Required if also registering the management cluster.")
	flags.BoolVar(&o.SkipChecks, "skip-checks", false, "If true, skip the pre-install checks.")

	cobra.MarkFlagRequired(flags, "license")
}

func (o EnterpriseOptions) getInstaller() *helm.Installer {
	ins := o.Options.getInstaller(gloomesh.GlooMeshEnterpriseChartUriTemplate)
	ins.ReleaseName = o.ReleaseName
	ins.Values["licenseKey"] = o.LicenseKey
	if o.SkipUI {
		ins.Values["gloo-mesh-ui.enabled"] = "false"
	}

	if o.IncludeRBAC {
		ins.Values["rbac-webhook.enabled"] = "true"
	} else {
		ins.Values["rbac-webhook.enabled"] = "false"
	}

	return &ins
}

func (o EnterpriseOptions) getRegistrationOptions() enterprise.RegistrationOptions {
	if o.RelayServerAddress == "" {
		namespacedLocalRelayServerAddress := fmt.Sprintf("enterprise-networking.%s.svc.cluster.local:9900", o.Namespace)
		logrus.Infof("No relay server address provided, defaulting to %s", namespacedLocalRelayServerAddress)
		o.RelayServerAddress = namespacedLocalRelayServerAddress
	}

	registrationOptions := enterprise.RegistrationOptions{
		Options:                o.Options.getRegistrationOptions(),
		AgentChartPathOverride: o.AgentChartPath,
		AgentChartValuesPath:   o.AgentValuesPath,
		RelayServerAddress:     o.RelayServerAddress,
		SkipChecks:             o.SkipChecks,
	}

	return registrationOptions
}

func InstallEnterprise(ctx context.Context, opts EnterpriseOptions) error {
	const (
		repoURI   = "https://storage.googleapis.com/gloo-mesh-enterprise"
		chartName = "gloo-mesh-enterprise"
	)
	if opts.Version == "" {
		cliVersion, err := semver.NewVersion(cliversion.Version)
		if err != nil {
			return eris.Wrapf(err, "invalid CLI version: %s", cliversion.Version)
		}
		stable := cliVersion.Prerelease() == "" // Get latest stable if not using a pre-release CLI
		version, err := helm.GetLatestChartMinorVersion(
			repoURI, chartName, stable,
			cliVersion.Major(), cliVersion.Minor(),
		)
		if err != nil {
			return err
		}
		opts.Version = version
	}

	if !opts.SkipChecks {
		logrus.Info("🔎 Performing server pre-install checks...")
		if err := runPreinstallChecks(ctx, &opts); err != nil {
			return eris.Wrap(err, "agent pre-install check failed")
		}
		logrus.Info("✅  server pre-install checks succeeded!")
	}

	// install relay server
	logrus.Info("💻 Installing relay server in the management cluster")
	if err := opts.getInstaller().InstallChart(ctx); err != nil {
		return eris.Wrap(err, "installing gloo-mesh-enterprise")
	}

	// install relay agent
	if opts.Register && !opts.DryRun {
		logrus.Info("Registering cluster")
		return enterprise.RegisterCluster(ctx, opts.getRegistrationOptions())
	}

	return nil
}

// run all server pre install checks
func runPreinstallChecks(ctx context.Context, opts *EnterpriseOptions) error {
	var crdMD map[string]*crdutils.CRDMetadata
	crdMDForDeploy, err := installutils.GetCrdMetadataFromInstaller(ctx, validationconsts.MgmtDeployName, opts.getInstaller())
	if err != nil {
		logrus.Warnf("Unable to fetch CRD metadata from Helm chart, unable to perform CRD upgrade check: %v", err)
		// we don't want to error if we can't get the CRD metadata, as this might be a release without one.
		// we may want to change that in the future.
	} else {
		crdMD = map[string]*crdutils.CRDMetadata{
			validationconsts.MgmtDeployName: crdMDForDeploy,
		}
	}

	checkCtx, err := validation.NewOutOfClusterCheckContext(
		opts.Options.KubeCfgPath,
		opts.Options.KubeContext,
		opts.Options.Namespace,
		0,
		0,
		nil,
		false,
		crdMD,
	)
	if err != nil {
		return err
	}

	if foundFailure := checks.RunChecks(ctx, checkCtx, checks.Server, checks.PreInstall); foundFailure {
		return eris.New("encountered failed pre-install checks")
	}

	return nil
}

func getStringMap(values []string) map[string]string {
	m := make(map[string]string)
	for _, e := range values {
		tokens := strings.Split(e, "=")
		k := strings.TrimSpace(tokens[0])
		v := strings.TrimSpace(tokens[1])
		m[k] = v
	}
	return m
}
