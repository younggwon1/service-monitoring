package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"

	k8s "github.com/younggwon1/service-monitoring/config/kubernetes"
	metrics "github.com/younggwon1/service-monitoring/config/metrics"
	pod "github.com/younggwon1/service-monitoring/pod"
)

// Swagger API Declaration
func setupSwagger(router *gin.Engine) {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Service Monitoring API"
	docs.SwaggerInfo.Description = "This is a service monitoring api using golang."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "0.0.0.0"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// use ginSwagger middleware to serve the API docs
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func main() {
	kubernetesConfig := k8s.GetKubernetesCredentials() // Init kubernetes config
	metricsConfig := metrics.GetMetricsCredentials()   // Init metrics config
	router := gin.Default()                            // Init Router
	setupSwagger(router)                               // Middleware Configuration
	// c := controller.NewController()                 // Init Controller

	v1 := router.Group("/v1")
	{
		podRouter := v1.Group("/podData")
		{
			podRouter.GET("", func(ctx *gin.Context) { pod.GetAllPodsData(ctx, kubernetesConfig) })
			podRouter.GET("/pod", func(ctx *gin.Context) { pod.GetSpecificPodData(ctx, kubernetesConfig) })
			podRouter.GET("/podLogs", func(ctx *gin.Context) { pod.GetSpecificPodLogs(ctx, kubernetesConfig) })
			podRouter.GET("/podMetrics", func(ctx *gin.Context) { pod.GetSpecificPodResourceUsage(ctx, metricsConfig) })
		}
	}
	router.Run(":8080")
}
