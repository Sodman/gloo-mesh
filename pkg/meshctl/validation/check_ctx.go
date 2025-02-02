package validation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/rotisserie/eris"
	apiextv1beta1 "github.com/solo-io/external-apis/pkg/api/k8s/apiextensions.k8s.io/v1beta1"
	appsv1 "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1"
	corev1 "github.com/solo-io/external-apis/pkg/api/k8s/core/v1"
	networkingv1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/common/defaults"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/utils"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/validation/checks"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/validation/consts"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"
	"github.com/solo-io/skv2/pkg/crdutils"
	skutils "github.com/solo-io/skv2/pkg/utils"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type InClusterCheckContext struct {
	checks.CommonContext
}

type OutOfClusterCheckContext struct {
	checks.CommonContext

	// this kubeconfig represents the mgmt cluster if running server check, and remote cluster if running agent check
	kubeConfig  string
	kubeContext string

	localPort  uint32
	remotePort uint32
}

func getCrdMetadataFromEnv() (map[string]*crdutils.CRDMetadata, error) {
	// We expect this variable to be set by the k8s downward api
	deploymentTested := os.Getenv(consts.DeploymentTestedDownwardApiEnvVar)
	if deploymentTested == "" {
		return nil, errors.New(consts.DeploymentTestedDownwardApiEnvVar + " env var not set")
	}
	crdJson := os.Getenv(consts.CrdMetadataDownwardApiEnvVar)
	if crdJson == "" {
		return nil, errors.New(consts.CrdMetadataDownwardApiEnvVar + " env var not set")
	}
	var crdMd crdutils.CRDMetadata
	err := json.Unmarshal([]byte(crdJson), &crdMd)
	if err != nil {
		return nil, err
	}
	return map[string]*crdutils.CRDMetadata{
		deploymentTested: &crdMd,
	}, nil
}

func NewInClusterCheckContext() (checks.CheckContext, error) {
	client, err := utils.BuildClient("", "")
	if err != nil {
		return nil, err
	}
	ns := os.Getenv("POD_NAMESPACE")
	if ns == "" {
		ns, err = skutils.GetInClusterNamesapce()
		if err != nil {
			return nil, err
		}
	}

	var skipChecks bool
	skipChecksEnv := os.Getenv("SKIP_CHECKS")
	if skipChecksEnv == "1" || strings.ToLower(skipChecksEnv) == "true" {
		skipChecks = true
	}

	// We expect this variable to be set by the k8s downward api
	crdMd, err := getCrdMetadataFromEnv()
	if err != nil {
		return nil, err
	}

	// get the crd annotations from the downward api:
	ret := &InClusterCheckContext{
		CommonContext: checks.CommonContext{
			Env: checks.Environment{
				AdminPort: defaults.MetricsPort,
				Namespace: ns,
				InCluster: true,
			},
			RelayDialer:             checks.NewRelayDialer(),
			AgentParams:             nil, // TODO pass in install / upgrade parameters, perhaps through CLI var args?
			CrdMetadata:             crdMd,
			SkipChecks:              skipChecks,
			AppsClientset:           appsv1.NewClientset(client),
			CoreClientset:           corev1.NewClientset(client),
			NetworkingClientset:     networkingv1.NewClientset(client),
			KubernetesClusterClient: v1alpha1.NewKubernetesClusterClient(client),
			CrdClient:               apiextv1beta1.NewCustomResourceDefinitionClient(client),
		},
	}

	return ret, nil
}

// exposed for testing, allows injecting mock k8s client
func NewTestCheckContext(
	gmInstallationNamespace string,
	localPort, remotePort uint32,
	agentParams *checks.AgentParams,
	relayDialer checks.RelayDialer,
	appsClientset appsv1.Clientset,
	coreClientset corev1.Clientset,
	networkingClientset networkingv1.Clientset,
	kubernetesClusterClient v1alpha1.KubernetesClusterClient,
	crdClient apiextv1beta1.CustomResourceDefinitionClient,
	ignoreChecks bool,
	crdMd map[string]*crdutils.CRDMetadata,
) checks.CheckContext {
	return &OutOfClusterCheckContext{
		remotePort: remotePort,
		localPort:  localPort,
		CommonContext: checks.CommonContext{
			Env: checks.Environment{
				AdminPort: remotePort,
				Namespace: gmInstallationNamespace,
				InCluster: false,
			},
			CrdMetadata:             crdMd,
			RelayDialer:             relayDialer,
			AgentParams:             agentParams,
			SkipChecks:              ignoreChecks,
			AppsClientset:           appsClientset,
			CoreClientset:           coreClientset,
			NetworkingClientset:     networkingClientset,
			KubernetesClusterClient: kubernetesClusterClient,
			CrdClient:               crdClient,
		},
	}
}

func NewOutOfClusterCheckContext(
	kubeConfig, kubeContext, gmInstallationNamespace string,
	localPort, remotePort uint32,
	agentParams *checks.AgentParams,
	ignoreChecks bool,
	crdMd map[string]*crdutils.CRDMetadata,
) (checks.CheckContext, error) {
	client, err := utils.BuildClient(kubeConfig, kubeContext)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to construct kube client from provided kubeconfig")
	}

	return &OutOfClusterCheckContext{
		remotePort:  remotePort,
		localPort:   localPort,
		kubeConfig:  kubeConfig,
		kubeContext: kubeContext,
		CommonContext: checks.CommonContext{
			Env: checks.Environment{
				AdminPort: remotePort,
				Namespace: gmInstallationNamespace,
				InCluster: false,
			},
			RelayDialer:             checks.NewRelayDialer(),
			AgentParams:             agentParams,
			CrdMetadata:             crdMd,
			SkipChecks:              ignoreChecks,
			AppsClientset:           appsv1.NewClientset(client),
			CoreClientset:           corev1.NewClientset(client),
			NetworkingClientset:     networkingv1.NewClientset(client),
			KubernetesClusterClient: v1alpha1.NewKubernetesClusterClient(client),
			CrdClient:               apiextv1beta1.NewCustomResourceDefinitionClient(client),
		},
	}, nil

}

func (c *InClusterCheckContext) Context() checks.CommonContext {
	return c.CommonContext
}

func (c *InClusterCheckContext) CRDMetadata(ctx context.Context, deploymentName string) (*crdutils.CRDMetadata, error) {
	crdMetadata, ok := c.CommonContext.CrdMetadata[deploymentName]
	if !ok {
		return nil, eris.Errorf("could not find CRD metadata for deployment name: %s", deploymentName)
	}
	return crdMetadata, nil
}

func (c *InClusterCheckContext) AccessAdminPort(ctx context.Context, deployment string, op func(ctx context.Context, adminUrl *url.URL) (error, string)) (error, string) {

	// note: the metrics port is not exposed on the service (it should not be, so this is fine).
	// so we need to find the ip of the deployed pod:
	d, err := c.Context().AppsClientset.Deployments().GetDeployment(ctx, client.ObjectKey{
		Namespace: c.Env.Namespace,
		Name:      deployment,
	})
	if err != nil {
		if kerrors.IsNotFound(err) {
			return err, "gloo-mesh enterprise deployment not found. Is gloo-mesh installed in this namespace?"
		}
		return err, ""
	}
	selector, err := metav1.LabelSelectorAsSelector(d.Spec.Selector)
	if err != nil {
		return err, ""
	}
	lo := &client.ListOptions{
		Namespace:     c.Env.Namespace,
		LabelSelector: selector,
		Limit:         1,
	}
	podsList, err := c.Context().CoreClientset.Pods().ListPod(ctx, lo)
	if err != nil {
		return err, "failed listing deployment pods. is gloo-mesh installed?"
	}
	pods := podsList.Items
	if len(pods) == 0 {
		return err, "no pods are available for deployemt. please check your gloo-mesh installation?"
	}
	if podsList.RemainingItemCount != nil && *podsList.RemainingItemCount != 0 {
		contextutils.LoggerFrom(ctx).Info("You have more than one pod for gloo-mesh deployment. This test may not be accurate.")
	}
	pod := pods[0]
	if pod.Status.PodIP == "" {
		return errors.New("no pod ip"), "gloo-mesh pod doesn't have an IP address. This is usually temporary. please wait or check your gloo-mesh installation?"
	}
	adminUrl := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%v:%v", pod.Status.PodIP, c.Env.AdminPort),
	}

	return op(ctx, adminUrl)
}

func (c *OutOfClusterCheckContext) Context() checks.CommonContext {
	return c.CommonContext
}

func (c *OutOfClusterCheckContext) AccessAdminPort(ctx context.Context, deployment string, op func(ctx context.Context, adminUrl *url.URL) (error, string)) (error, string) {
	portFwdContext, cancelPtFwd := context.WithCancel(ctx)
	defer cancelPtFwd()

	// start port forward to mgmt server stats port
	localPort, err := utils.PortForwardFromDeployment(
		portFwdContext,
		c.kubeConfig,
		c.kubeContext,
		deployment,
		c.Env.Namespace,
		fmt.Sprintf("%v", c.localPort),
		fmt.Sprintf("%v", c.remotePort),
	)
	if err != nil {
		return err, fmt.Sprintf("try verifying that `kubectl port-forward -n %v deployment/%v %v:%v` can be run successfully.", c.Env.Namespace, deployment, c.localPort, c.remotePort)
	}
	// request metrics page from mgmt deployment
	adminUrl := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("localhost:%v", localPort),
	}

	return op(portFwdContext, adminUrl)
}

// TODO We need a cleaner way to signal whether CRD metadata should be fetched from the Deployment on the cluster, which
// should only happen in the post-install check. For pre-install checks, CRD metadata is expected to be passed in via the check constructor,
// which, with the current logic, requires us to express missing CRD metadata (e.g. for older versions of GM) via a non-nil map containing a key for the
// deployment with a nil value.
func (c *OutOfClusterCheckContext) CRDMetadata(ctx context.Context, deploymentName string) (*crdutils.CRDMetadata, error) {
	if c.CommonContext.CrdMetadata != nil {
		if md, ok := c.CommonContext.CrdMetadata[deploymentName]; ok {
			return md, nil
		}
	}
	// if not provided in construction, fetch from the cluster:
	// read the annotations of the deployment
	d, err := c.AppsClientset.Deployments().GetDeployment(ctx, client.ObjectKey{
		Namespace: c.Env.Namespace,
		Name:      deploymentName,
	})
	if err != nil {
		return nil, err
	}

	crdMeta, err := crdutils.ParseCRDMetadataFromAnnotations(d.Annotations)
	if err != nil {
		return nil, err
	}
	if crdMeta == nil {
		return nil, eris.Errorf("No CRD metadata found for deployment %s", deploymentName)
	}

	if c.CommonContext.CrdMetadata == nil {
		c.CommonContext.CrdMetadata = make(map[string]*crdutils.CRDMetadata)
	}

	c.CommonContext.CrdMetadata[deploymentName] = crdMeta
	return crdMeta, nil
}
