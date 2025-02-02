package utils

import (
	"context"

	v1 "github.com/solo-io/external-apis/pkg/api/k8s/core/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func EnsureNamespace(ctx context.Context, kubeClient client.Client, namespace string) error {
	namespaces := v1.NewNamespaceClient(kubeClient)
	return namespaces.UpsertNamespace(ctx, &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:   namespace,
			Labels: map[string]string{"istio-injection": "disabled"},
		},
		Spec: corev1.NamespaceSpec{Finalizers: []corev1.FinalizerName{"kubernetes"}},
	})
}

func CleanupCompletedPodsInNamespace(ctx context.Context, kubeconfig, kubecontext, namespace string) error {
	kubeClient, err := BuildClient(kubeconfig, kubecontext)
	if err != nil {
		return err
	}
	pods := v1.NewPodClient(kubeClient)

	return pods.DeleteAllOfPod(ctx, client.InNamespace(namespace), client.MatchingFieldsSelector{
		Selector: fields.OneTermEqualSelector("status.phase", "Succeeded"),
	})
}
