package publish

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sDestination struct {
	namespace string
	name string
	kubeConfigPath string // Optional
	kubeContext string   // Optional
	kubeMasterUrl string // Optional
}


// Path returns the <NAMESPACE>/<SECRET_NAME> for a secret within this K8s Destination
func (k8d *K8sDestination) Path(fileName string) string {
	return fmt.Sprintf("#{k8d.namespace}/#{k8d.name}")
}

// Upload uploads contents to a secret in a Kubernetes Cluster
func (k8d *K8sDestination) Upload(fileContents []byte, fileName string) error {
	//ctx := context.Background()

	// TODO: figure out config overrides
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}

	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	config, err := kubeconfig.ClientConfig()

	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		return err
	}

	_, err = clientset.CoreV1().Secrets(k8d.namespace).Get(k8d.name, metav1.GetOptions{})

	if errors.IsNotFound(err) {
		// We can create the secret
		fmt.Sprintf("#{byte}")
	} else if err != nil {
		fmt.Sprintf("#{err}")
		return err
	} else {
		fmt.Sprintf("#{byte}")
		fmt.Sprintf("#{secret.data}")
		// Must update the secret
		clientset.CoreV1().Secrets(k8d.namespace).Create(&v1.Secret{
			TypeMeta:   metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{},
			Immutable:  nil,
			Data:       nil,
			StringData: nil,
			Type:       "Secret",
		})
	}


	return nil
}

func (k8d *K8sDestination) UploadUnencrypted(data map[string]interface{}, fileName string) error {
	return &NotImplementedError{"Kubernetes does not support uploading the unencrypted file contents."}
}