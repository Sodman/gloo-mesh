package mtls_test

import (
	"context"
	"time"

	commonv1 "github.com/solo-io/gloo-mesh/pkg/api/common.mesh.gloo.solo.io/v1"
	"istio.io/istio/security/pkg/pki/util"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"
	certificatesv1 "github.com/solo-io/gloo-mesh/pkg/api/certificates.mesh.gloo.solo.io/v1"
	discoveryv1 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1"
	discoveryv1sets "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1/sets"
	mock_istio "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/istio/mocks"
	mock_local "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/output/local/mocks"
	networkingv1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"
	"github.com/solo-io/gloo-mesh/pkg/certificates/common/secrets"
	mock_reporting "github.com/solo-io/gloo-mesh/pkg/mesh-networking/reporting/mocks"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/istio/mesh/mtls"
	"github.com/solo-io/gloo-mesh/pkg/mesh-networking/translation/utils/metautils"
	skv2corev1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/ezkube"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("MtlsTranslator", func() {
	var (
		ctrl *gomock.Controller
		ctx  context.Context

		mockIstioBuilder *mock_istio.MockBuilder
		mockLocalBuilder *mock_local.MockBuilder
		mockReporter     *mock_reporting.MockReporter

		istioMesh          *discoveryv1.Mesh
		ingressDestination *discoveryv1.Destination
		childResourceMeta  *metav1.ObjectMeta
	)

	BeforeEach(func() {
		ctrl, ctx = gomock.WithContext(context.Background(), GinkgoT())
		mockIstioBuilder = mock_istio.NewMockBuilder(ctrl)
		mockLocalBuilder = mock_local.NewMockBuilder(ctrl)
		mockReporter = mock_reporting.NewMockReporter(ctrl)

		ingressDestination = &discoveryv1.Destination{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "istio-ingressgateway",
				Namespace: "istio-system",
			},
			Spec: discoveryv1.DestinationSpec{
				Type: &discoveryv1.DestinationSpec_KubeService_{
					KubeService: &discoveryv1.DestinationSpec_KubeService{
						Ref: &skv2corev1.ClusterObjectRef{
							Name:        "istio-ingressgateway",
							Namespace:   "istio-system",
							ClusterName: "cluster-name",
						},
						WorkloadSelectorLabels: map[string]string{"gateway": "selector"},
					},
				},
			},
		}

		istioMesh = &discoveryv1.Mesh{
			ObjectMeta: metav1.ObjectMeta{
				Name: "my-istio-mesh",
			},
			Spec: discoveryv1.MeshSpec{
				Type: &discoveryv1.MeshSpec_Istio_{
					Istio: &discoveryv1.MeshSpec_Istio{
						TrustDomain:          "cluster.not-local",
						IstiodServiceAccount: "istiod-not-standard",
						Installation: &discoveryv1.MeshInstallation{
							Namespace: "istio-system-2",
							Cluster:   "cluster-name",
						},
					},
				},
				AgentInfo: &discoveryv1.MeshSpec_AgentInfo{
					AgentNamespace: "gloo-mesh",
				},
			},
			Status: discoveryv1.MeshStatus{
				AppliedEastWestIngressGateways: []*commonv1.AppliedIngressGateway{
					{
						DestinationRef: ezkube.MakeObjectRef(ingressDestination),
					},
				},
			},
		}

		childResourceMeta = &metav1.ObjectMeta{
			Name:        istioMesh.GetName(),
			Namespace:   istioMesh.Spec.GetAgentInfo().GetAgentNamespace(),
			ClusterName: istioMesh.Spec.GetIstio().GetInstallation().GetCluster(),
			Labels:      metautils.TranslatedObjectLabels(),
		}

	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("will skip if non-istio mesh", func() {
		translator := mtls.NewTranslator(ctx, nil, nil)
		mesh := &discoveryv1.Mesh{}
		vm := &discoveryv1.MeshStatus_AppliedVirtualMesh{}
		translator.Translate(mesh, vm, mockIstioBuilder, mockLocalBuilder, mockReporter)
	})

	It("generated root CA", func() {
		certSettings := &certificatesv1.CommonCertOptions{
			OrgName: "my-org",
		}

		vm := &discoveryv1.MeshStatus_AppliedVirtualMesh{
			Ref: &skv2corev1.ObjectRef{
				Name:      "my-vm",
				Namespace: "gloo-mesh",
			},
			Spec: &networkingv1.VirtualMeshSpec{
				MtlsConfig: &networkingv1.VirtualMeshSpec_MTLSConfig{
					TrustModel: &networkingv1.VirtualMeshSpec_MTLSConfig_Shared{
						Shared: &networkingv1.SharedTrust{
							CertificateAuthority: &networkingv1.SharedTrust_RootCertificateAuthority{
								RootCertificateAuthority: &networkingv1.RootCertificateAuthority{
									CaSource: &networkingv1.RootCertificateAuthority_Generated{
										Generated: certSettings,
									},
								},
							},
						},
					},
				},
			},
		}

		// Can't actually check against the exact context
		mockLocalBuilder.EXPECT().AddSecrets(gomock.Any()).Do(func(secret *corev1.Secret) {
			// Make sure the name is the same, maybe decode the cert and check the data?
			Expect(secret.GetName()).To(Equal(vm.GetRef().GetName() + "." + vm.GetRef().GetNamespace()))
			certData := secrets.CADataFromSecretData(secret.Data)
			err := certData.Verify()
			Expect(err).NotTo(HaveOccurred())
		})

		mockIstioBuilder.EXPECT().
			AddIssuedCertificates(gomock.Any()).
			Do(func(issuedCert *certificatesv1.IssuedCertificate) {
				cert := &certificatesv1.IssuedCertificate{
					ObjectMeta: *childResourceMeta,
					Spec: certificatesv1.IssuedCertificateSpec{
						Hosts: []string{"spiffe://cluster.not-local/ns/istio-system-2/sa/istiod-not-standard"},
						Org:   "Istio",
						CertOptions: &certificatesv1.CommonCertOptions{
							TtlDays:                        365,
							RsaKeySizeBytes:                4096,
							OrgName:                        "Istio",
							SecretRotationGracePeriodRatio: 0.10,
						},
						SigningCertificateSecret: &skv2corev1.ObjectRef{
							Name:      vm.GetRef().GetName() + "." + vm.GetRef().GetNamespace(),
							Namespace: "gloo-mesh",
						},
						CertificateAuthority: &certificatesv1.IssuedCertificateSpec_GlooMeshCa{
							GlooMeshCa: &certificatesv1.RootCertificateAuthority{
								CertificateAuthority: &certificatesv1.RootCertificateAuthority_SigningCertificateSecret{
									SigningCertificateSecret: &skv2corev1.ObjectRef{
										Name:      vm.GetRef().GetName() + "." + vm.GetRef().GetNamespace(),
										Namespace: "gloo-mesh",
									},
								},
							},
						},
						IssuedCertificateSecret: &skv2corev1.ObjectRef{
							Name:      "cacerts",
							Namespace: istioMesh.Spec.GetIstio().GetInstallation().GetNamespace(),
						},
					},
				}
				metautils.AppendParent(ctx, cert, vm.GetRef(), networkingv1.VirtualMesh{}.GVK())
				Expect(cert).To(Equal(issuedCert))
			})

		mockIstioBuilder.EXPECT().AddPodBounceDirectives(nil)

		translator := mtls.NewTranslator(ctx, v1sets.NewSecretSet(), nil)

		translator.Translate(istioMesh, vm, mockIstioBuilder, mockLocalBuilder, mockReporter)
	})

	It("provided root CA", func() {
		// Generate cert for the provided secret
		cert, key, err := util.GenCertKeyFromOptions(util.CertOptions{
			RSAKeySize:   2048,
			IsSelfSigned: true,
			IsCA:         true,
			TTL:          time.Minute,
		})
		Expect(err).NotTo(HaveOccurred())

		rootCaData := &secrets.CAData{
			CaPrivateKey: key,
			CaCert:       cert,
			RootCert:     cert,
			CertChain:    cert,
		}

		generatedSecret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-secret",
				Namespace: "my-namespace",
			},
			Data: rootCaData.ToSecretData(),
		}

		vm := &discoveryv1.MeshStatus_AppliedVirtualMesh{
			Ref: &skv2corev1.ObjectRef{
				Name:      "my-vm",
				Namespace: "gloo-mesh",
			},
			Spec: &networkingv1.VirtualMeshSpec{
				MtlsConfig: &networkingv1.VirtualMeshSpec_MTLSConfig{
					AutoRestartPods: true,
					TrustModel: &networkingv1.VirtualMeshSpec_MTLSConfig_Shared{
						Shared: &networkingv1.SharedTrust{
							CertificateAuthority: &networkingv1.SharedTrust_RootCertificateAuthority{
								RootCertificateAuthority: &networkingv1.RootCertificateAuthority{
									CaSource: &networkingv1.RootCertificateAuthority_Secret{
										Secret: &skv2corev1.ObjectRef{
											Name:      generatedSecret.GetName(),
											Namespace: generatedSecret.GetNamespace(),
										},
									},
								},
							},
						},
					},
				},
			},
		}

		workload := &discoveryv1.Workload{
			Spec: discoveryv1.WorkloadSpec{
				Mesh: ezkube.MakeObjectRef(istioMesh),
				Type: &discoveryv1.WorkloadSpec_Kubernetes{
					Kubernetes: &discoveryv1.WorkloadSpec_KubernetesWorkload{
						Controller: &skv2corev1.ClusterObjectRef{
							Namespace: "namespace",
						},
						PodLabels: map[string]string{
							"pod": "labels",
						},
					},
				},
			},
		}

		expectedPbd := &certificatesv1.PodBounceDirective{
			ObjectMeta: metav1.ObjectMeta{
				Name:        istioMesh.GetName(),
				Namespace:   istioMesh.Spec.AgentInfo.GetAgentNamespace(),
				Labels:      metautils.TranslatedObjectLabels(),
				ClusterName: istioMesh.Spec.GetIstio().GetInstallation().GetCluster(),
			},
			Spec: certificatesv1.PodBounceDirectiveSpec{
				PodsToBounce: []*certificatesv1.PodBounceDirectiveSpec_PodSelector{
					{
						Namespace:       istioMesh.Spec.GetIstio().Installation.GetNamespace(),
						Labels:          istioMesh.Spec.GetIstio().Installation.GetPodLabels(),
						WaitForReplicas: 1,
					},
					{
						Namespace: workload.Spec.GetKubernetes().GetController().GetNamespace(),
						Labels:    workload.Spec.GetKubernetes().GetPodLabels(),
						RootCertSync: &certificatesv1.PodBounceDirectiveSpec_PodSelector_RootCertSync{
							SecretRef: &skv2corev1.ObjectRef{
								Name:      "cacerts",
								Namespace: istioMesh.Spec.GetIstio().GetInstallation().GetNamespace(),
							},
							SecretKey:    secrets.RootCertID,
							ConfigMapKey: secrets.RootCertID,
							ConfigMapRef: &skv2corev1.ObjectRef{
								Name:      "istio-ca-root-cert",
								Namespace: workload.Spec.GetKubernetes().GetController().GetNamespace(),
							},
						},
					},
				},
			},
		}
		metautils.AppendParent(ctx, expectedPbd, vm.GetRef(), networkingv1.VirtualMesh{}.GVK())

		mockIstioBuilder.EXPECT().
			AddIssuedCertificates(gomock.Any()).
			Do(func(issuedCert *certificatesv1.IssuedCertificate) {
				cert := &certificatesv1.IssuedCertificate{
					ObjectMeta: *childResourceMeta,
					Spec: certificatesv1.IssuedCertificateSpec{
						Hosts: []string{"spiffe://cluster.not-local/ns/istio-system-2/sa/istiod-not-standard"},
						Org:   "Istio",
						SigningCertificateSecret: &skv2corev1.ObjectRef{
							Name:      generatedSecret.GetName(),
							Namespace: generatedSecret.GetNamespace(),
						},
						CertOptions: &certificatesv1.CommonCertOptions{
							TtlDays:                        365,
							RsaKeySizeBytes:                4096,
							OrgName:                        "Istio",
							SecretRotationGracePeriodRatio: 0.10,
						},
						CertificateAuthority: &certificatesv1.IssuedCertificateSpec_GlooMeshCa{
							GlooMeshCa: &certificatesv1.RootCertificateAuthority{
								CertificateAuthority: &certificatesv1.RootCertificateAuthority_SigningCertificateSecret{
									SigningCertificateSecret: &skv2corev1.ObjectRef{
										Name:      generatedSecret.GetName(),
										Namespace: generatedSecret.GetNamespace(),
									},
								},
							},
						},
						PodBounceDirective: ezkube.MakeObjectRef(expectedPbd),
						IssuedCertificateSecret: &skv2corev1.ObjectRef{
							Name:      "cacerts",
							Namespace: istioMesh.Spec.GetIstio().GetInstallation().GetNamespace(),
						},
					},
				}
				metautils.AppendParent(ctx, cert, vm.GetRef(), networkingv1.VirtualMesh{}.GVK())
				Expect(cert).To(Equal(issuedCert))
			})

		mockIstioBuilder.EXPECT().
			AddPodBounceDirectives(gomock.Any()).
			Do(func(podBounceDirective *certificatesv1.PodBounceDirective) {
				Expect(podBounceDirective).To(Equal(expectedPbd))
			})

		translator := mtls.NewTranslator(
			ctx,
			v1sets.NewSecretSet(generatedSecret),
			discoveryv1sets.NewWorkloadSet(workload),
		)

		translator.Translate(istioMesh, vm, mockIstioBuilder, mockLocalBuilder, mockReporter)
	})

	It("Intermediate CA", func() {

		intermediateCa := &certificatesv1.IntermediateCertificateAuthority{
			CaSource: &certificatesv1.IntermediateCertificateAuthority_Vault{
				Vault: &certificatesv1.VaultCA{
					CaPath: "path-to-ca",
				},
			},
		}

		vm := &discoveryv1.MeshStatus_AppliedVirtualMesh{
			Ref: &skv2corev1.ObjectRef{
				Name:      "my-vm",
				Namespace: "gloo-mesh",
			},
			Spec: &networkingv1.VirtualMeshSpec{
				MtlsConfig: &networkingv1.VirtualMeshSpec_MTLSConfig{
					AutoRestartPods: true,
					TrustModel: &networkingv1.VirtualMeshSpec_MTLSConfig_Shared{
						Shared: &networkingv1.SharedTrust{
							CertificateAuthority: &networkingv1.SharedTrust_IntermediateCertificateAuthority{
								IntermediateCertificateAuthority: &certificatesv1.IntermediateCertificateAuthority{
									CaSource: &certificatesv1.IntermediateCertificateAuthority_Vault{
										Vault: &certificatesv1.VaultCA{
											CaPath: "path-to-ca",
										},
									},
								},
							},
						},
					},
				},
			},
		}

		kubeWorkload := &discoveryv1.Workload{
			Spec: discoveryv1.WorkloadSpec{
				Mesh: ezkube.MakeObjectRef(istioMesh),
				Type: &discoveryv1.WorkloadSpec_Kubernetes{
					Kubernetes: &discoveryv1.WorkloadSpec_KubernetesWorkload{
						Controller: &skv2corev1.ClusterObjectRef{
							Namespace: "namespace",
						},
						PodLabels: map[string]string{
							"pod": "labels",
						},
					},
				},
			},
		}

		pbd := &certificatesv1.PodBounceDirective{
			ObjectMeta: metav1.ObjectMeta{
				Name:        istioMesh.GetName(),
				Namespace:   istioMesh.Spec.AgentInfo.GetAgentNamespace(),
				Labels:      metautils.TranslatedObjectLabels(),
				ClusterName: istioMesh.Spec.GetIstio().GetInstallation().GetCluster(),
			},
			Spec: certificatesv1.PodBounceDirectiveSpec{
				PodsToBounce: []*certificatesv1.PodBounceDirectiveSpec_PodSelector{
					{
						Namespace: kubeWorkload.Spec.GetKubernetes().GetController().GetNamespace(),
						Labels:    kubeWorkload.Spec.GetKubernetes().GetPodLabels(),
						RootCertSync: &certificatesv1.PodBounceDirectiveSpec_PodSelector_RootCertSync{
							SecretRef: &skv2corev1.ObjectRef{
								Name:      "cacerts",
								Namespace: istioMesh.Spec.GetIstio().GetInstallation().GetNamespace(),
							},
							ConfigMapRef: &skv2corev1.ObjectRef{
								Name:      "istio-ca-root-cert",
								Namespace: kubeWorkload.Spec.GetKubernetes().GetController().GetNamespace(),
							},
							SecretKey:    secrets.RootCertID,
							ConfigMapKey: secrets.RootCertID,
						},
					},
				},
			},
		}
		metautils.AppendParent(ctx, pbd, vm.GetRef(), networkingv1.VirtualMesh{}.GVK())

		mockIstioBuilder.EXPECT().
			AddIssuedCertificates(gomock.Any()).
			Do(func(issuedCert *certificatesv1.IssuedCertificate) {
				cert := &certificatesv1.IssuedCertificate{
					ObjectMeta: *childResourceMeta,
					Spec: certificatesv1.IssuedCertificateSpec{
						Hosts: []string{"spiffe://cluster.not-local/ns/istio-system-2/sa/istiod-not-standard"},
						Org:   "Istio",
						CertOptions: &certificatesv1.CommonCertOptions{
							TtlDays:                        365,
							RsaKeySizeBytes:                4096,
							OrgName:                        "Istio",
							SecretRotationGracePeriodRatio: 0.10,
						},
						CertificateAuthority: &certificatesv1.IssuedCertificateSpec_AgentCa{
							AgentCa: intermediateCa,
						},
						PodBounceDirective: ezkube.MakeObjectRef(pbd),
					},
				}
				metautils.AppendParent(ctx, cert, vm.GetRef(), networkingv1.VirtualMesh{}.GVK())
				Expect(cert).To(Equal(issuedCert))
			})

		mockIstioBuilder.EXPECT().
			AddPodBounceDirectives(gomock.Any()).
			Do(func(podBounceDirective *certificatesv1.PodBounceDirective) {
				Expect(pbd).To(Equal(podBounceDirective))
			})

		translator := mtls.NewTranslator(ctx, nil, discoveryv1sets.NewWorkloadSet(kubeWorkload))

		translator.Translate(istioMesh, vm, mockIstioBuilder, mockLocalBuilder, mockReporter)
	})

})
