package watch

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

func GetAllDeploymentsData(ctx *gin.Context, clientSet *kubernetes.Clientset) {
	deploymentWatch, err := clientSet.AppsV1().Deployments("").Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for event := range deploymentWatch.ResultChan() {
		item := event.Object.(*v1.Deployment)

		switch event.Type {
		case watch.Modified:
			continue
		case watch.Bookmark:
			continue
		case watch.Error:
			fmt.Println(item.GetName())
		case watch.Deleted:
			continue
			// fmt.Println(item.GetName())
		case watch.Added:
			continue
			// fmt.Println(item.GetName())
			// processNamespace(item.GetName())
		}
	}
}
