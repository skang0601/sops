package publish

import (
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"
)

type KubernetesDestination struct {
	namespace string
	name string
}

func NewKubernetesDestination(namespace, name string) *KubernetesDestination {
	return &KubernetesDestination{namespace, name}
}
