package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func GetKubernetesCredentials() *kubernetes.Clientset {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err.Error())
	}

	return clientSet
}
