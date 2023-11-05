package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"

	k8s "github.com/younggwon1/service-monitoring/config/kubernetes"
	pod "github.com/younggwon1/service-monitoring/pod"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func getSpecificPodLogs(c *gin.Context) {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err.Error())
	}
	podLogOpts := corev1.PodLogOptions{}

	nameSpace := c.DefaultQuery("nameSpace", "default")
	podName := c.Query("podName")

	req := clientSet.CoreV1().Pods(nameSpace).GetLogs(podName, &podLogOpts)

	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		panic(err.Error())
	}

	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		panic(err.Error())
	}

	str := buf.String()

	fmt.Println(str) // 확인 필요.

}

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
	router := gin.Default()                            // Init Router
	setupSwagger(router)                               // Middleware Configuration
	// c := controller.NewController()                 // Init Controller

	v1 := router.Group("/v1")
	{
		podRouter := v1.Group("/podData")
		{
			podRouter.GET("/pods", func(ctx *gin.Context) { pod.GetAllPodStatus(ctx, kubernetesConfig) })
			podRouter.GET("/pod", func(ctx *gin.Context) { pod.GetSpecificPodStatus(ctx, kubernetesConfig) })
			podRouter.GET("/podLogs", getSpecificPodLogs)
		}
	}
	router.Run(":8080")
}
