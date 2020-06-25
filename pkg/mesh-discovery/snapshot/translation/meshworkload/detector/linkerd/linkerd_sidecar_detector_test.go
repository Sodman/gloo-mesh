package linkerd_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	v1alpha1sets "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1/sets"
	. "github.com/solo-io/service-mesh-hub/pkg/mesh-discovery/snapshot/translation/meshworkload/detector/linkerd"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("LinkerdSidecarDetector", func() {
	serviceAccountName := "service-account-name"
	ns := "namespace"
	clusterName := "cluster"
	podName := "pod"

	pod := func() *corev1.Pod {
		return &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:   ns,
				Name:        podName,
				ClusterName: clusterName,
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name: "linkerd-proxy",
					},
				},
				ServiceAccountName: serviceAccountName,
			},
		}
	}

	linkerdMeshes := func(cluster string) v1alpha1sets.MeshSet {
		return v1alpha1sets.NewMeshSet(
			&v1alpha1.Mesh{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "linkerd-system",
					Name:      "linkerd-cluster",
				},
				Spec: v1alpha1.MeshSpec{
					MeshType: &v1alpha1.MeshSpec_Linkerd{
						Linkerd: &v1alpha1.MeshSpec_LinkerdMesh{
							Installation: &v1alpha1.MeshSpec_MeshInstallation{
								Cluster: cluster,
							},
						},
					},
				},
			},
		)
	}

	detector := NewSidecarDetector(context.TODO())

	It("detects workload when sidecar mesh is in cluster", func() {
		pod := pod()

		meshes := linkerdMeshes(clusterName)

		meshWorkload := detector.DetectMeshSidecar(pod, meshes)
		Expect(meshWorkload).To(Equal(meshes.List()[0]))
	})
	It("does not detect workload when sidecar mesh is of different cluster", func() {
		pod := pod()

		meshes := linkerdMeshes("different-" + clusterName)

		meshWorkload := detector.DetectMeshSidecar(pod, meshes)
		Expect(meshWorkload).To(BeNil())
	})
	It("does not detect workload when sidecar mesh is not present", func() {
		pod := pod()

		meshes := v1alpha1sets.NewMeshSet()

		meshWorkload := detector.DetectMeshSidecar(pod, meshes)
		Expect(meshWorkload).To(BeNil())
	})
	It("does not detect workload when sidecar is not present", func() {
		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:   ns,
				Name:        podName,
				ClusterName: clusterName,
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Image: "blah",
					},
				},
			},
		}

		meshes := linkerdMeshes(clusterName)

		meshWorkload := detector.DetectMeshSidecar(pod, meshes)
		Expect(meshWorkload).To(BeNil())
	})

})
