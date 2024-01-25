package deployment

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	// request "github.com/younggwon1/service-monitoring/deployment/config/request"
	response "github.com/younggwon1/service-monitoring/deployment/config/response"
)

func AllDeployments(ctx *gin.Context, clientSet *kubernetes.Clientset) {
	deployments, err := clientSet.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var deploymentsList []response.ResponseAllDeployments
	for _, deployment := range deployments.Items {
		deploymentsInfo := &response.ResponseAllDeployments{
			Name:        deployment.Name,
			NameSpace:   deployment.Namespace,
			Age:         deployment.ObjectMeta.GetCreationTimestamp().Format("2006-01-02 15:04:05"),
			Current:     deployment.Status.Replicas,
			Desired:     deployment.Status.ReadyReplicas,
			UptoDate:    deployment.Status.UpdatedReplicas,
			Available:   deployment.Status.AvailableReplicas,
			UnAvailable: deployment.Status.UnavailableReplicas,
		}
		deploymentsList = append(deploymentsList, *deploymentsInfo)
	}

	if len(deploymentsList) != 0 {
		deploymentsListJson, _ := json.Marshal(deploymentsList)
		ctx.JSON(http.StatusOK, string(deploymentsListJson))
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "There is no deployment."})
	}
}
