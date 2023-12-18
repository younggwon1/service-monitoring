package metrics

import (
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func GetMetricsCredentials() *metrics.Clientset {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err.Error())
	}

	metricsCredentials, err := metrics.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	return metricsCredentials
}
