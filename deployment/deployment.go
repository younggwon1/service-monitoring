package deployment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	// request "github.com/younggwon1/service-monitoring/deployment/config/request"
	response "github.com/younggwon1/service-monitoring/deployment/config/response"
)

func ErrorDeployments(ctx *gin.Context, clientSet *kubernetes.Clientset) {
	deployments, err := clientSet.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var errorDeploymentsList []response.ResponseErrorDeployments
	for _, deployment := range deployments.Items {
		// Check the deployment error status
		if deployment.Status.Replicas != deployment.Status.AvailableReplicas {
			errorDeployment := &response.ResponseErrorDeployments{
				Name:      deployment.Name,
				NameSpace: deployment.Namespace,
				Reason:    string(deployment.Status.Conditions[1].Reason),
				Message:   deployment.Status.Conditions[1].Message,
				Age:       deployment.ObjectMeta.GetCreationTimestamp().Format("2006-01-02 15:04:05"),
			}
			errorDeploymentsList = append(errorDeploymentsList, *errorDeployment)
		}
	}

	if len(errorDeploymentsList) != 0 {
		errorDeploymentsJson, _ := json.Marshal(errorDeploymentsList)
		ctx.JSON(http.StatusOK, string(errorDeploymentsJson))
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "There is no deployment status as error."})
	}
}

func DeleteErrorDeployments(ctx *gin.Context, clientSet *kubernetes.Clientset) {
	fmt.Println("Delete Error Deployments")
}
