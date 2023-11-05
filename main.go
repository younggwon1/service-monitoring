package pod

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type requestSpecificPodData struct {
	Name      string `form:"name"`
	NameSpace string `form:"namespace"`
}

type responseAllPodData struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	Cluster   string `json:"cluster"`
	NameSpace string `json:"namespace"`
	Age       string `json:"age"`
	Restarts  string `json:"restarts"`
}

type responseSpecificPodData struct {
	Name      string                        `json:"name"`
	Namespace string                        `json:"namespace"`
	Status    responseSpecificPodStatusData `json:"status"`
	Spec      responseSpecificPodSpecData   `json:"spec"`
	Meta      responseSpecificPodMetaData   `json:"meta"`
}

type responseSpecificPodStatusData struct {
	InitContainerStatuses []v1.ContainerStatus `json:"initContainerStatuses"`
	ContainerStatuses     []v1.ContainerStatus `json:"containerStatuses"`
	HostIP                string               `json:"hostIP"`
	PodIP                 string               `json:"podIP"`
	Phase                 v1.PodPhase          `json:"phase"`
	QOSClass              v1.PodQOSClass       `json:"qosClass"`
	Message               string               `json:"message"`
}

type responseSpecificPodSpecData struct {
	InitContainerStatuses []v1.Container         `json:"initContainerStatuses"`
	ContainerStatuses     []v1.Container         `json:"containerStatuses"`
	RestartPolicy         v1.RestartPolicy       `json:"restartPolicy"`
	SchedulerName         string                 `json:"schedulerName"`
	SecurityContext       *v1.PodSecurityContext `json:"securityContext"`
	ServiceAccountName    string                 `json:"serviceAccountName"`
	Volumes               []v1.Volume            `json:"volumes"`
}

type responseSpecificPodMetaData struct {
	Name              string                  `json:"name"`
	Namespace         string                  `json:"namespace"`
	Labels            map[string]string       `json:"labels"`
	Annotations       map[string]string       `json:"annotations"`
	CreationTimestamp metav1.Time             `json:"creationTimestamp"`
	DeletionTimestamp *metav1.Time            `json:"deletionTimestamp"`
	OwnerReferences   []metav1.OwnerReference `json:"ownerReferences"`
}

func GetAllPodStatus(ctx *gin.Context, clientSet *kubernetes.Clientset) {
	getPodData, err := clientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// AS-IS
	// getPodDataJson, _ := json.Marshal(getPodData)
	// ctx.JSON(http.StatusOK, string(getPodDataJson))

	// TO-BE
	// Name      string `json:"name"`
	// Status    string `json:"status"`
	// Cluster   string `json:"cluster"`
	// NameSpace string `json:"namespace"`
	// Age       string `json:"age"`
	// Restarts  string `json:"restarts"`
	fmt.Println(getPodData.Items[0].Name)
	fmt.Println(getPodData.Items[0].Status.Conditions)
	fmt.Println(getPodData.Items[0].Status.StartTime)
	fmt.Println(getPodData.Items[0].Namespace)
	// fmt.Println(getPodData.Items[0])
	// fmt.Println(getPodData.Items[0])

	customGetAllPodData := &responseAllPodData{
		Name: getPodData.Continue,
	}

	customGetPodDataJson, _ := json.Marshal(customGetAllPodData)
	ctx.JSON(http.StatusOK, string(customGetPodDataJson))
}

func GetSpecificPodStatus(ctx *gin.Context, clientSet *kubernetes.Clientset) {
	var RequestSpecificPodData requestSpecificPodData

	if err := ctx.ShouldBindQuery(&RequestSpecificPodData); err == nil {
		getSpecificPodData, err := clientSet.CoreV1().Pods(RequestSpecificPodData.NameSpace).Get(context.TODO(), RequestSpecificPodData.Name, metav1.GetOptions{})
		if err != nil {
			panic(err.Error())
		}

		customGetSpecificPodData := &responseSpecificPodData{
			Name:      getSpecificPodData.Name,
			Namespace: getSpecificPodData.Namespace,
			Status: responseSpecificPodStatusData{
				InitContainerStatuses: getSpecificPodData.Status.InitContainerStatuses,
				ContainerStatuses:     getSpecificPodData.Status.InitContainerStatuses,
				HostIP:                getSpecificPodData.Status.HostIP,
				PodIP:                 getSpecificPodData.Status.PodIP,
				Phase:                 getSpecificPodData.Status.Phase,
				QOSClass:              getSpecificPodData.Status.QOSClass,
				Message:               getSpecificPodData.Status.Message,
			},
			Spec: responseSpecificPodSpecData{
				InitContainerStatuses: getSpecificPodData.Spec.InitContainers,
				ContainerStatuses:     getSpecificPodData.Spec.Containers,
				RestartPolicy:         getSpecificPodData.Spec.RestartPolicy,
				SchedulerName:         getSpecificPodData.Spec.SchedulerName,
				SecurityContext:       getSpecificPodData.Spec.SecurityContext,
				ServiceAccountName:    getSpecificPodData.Spec.ServiceAccountName,
				Volumes:               getSpecificPodData.Spec.Volumes,
			},
			Meta: responseSpecificPodMetaData{
				Name:              getSpecificPodData.ObjectMeta.Name,
				Namespace:         getSpecificPodData.ObjectMeta.Namespace,
				Labels:            getSpecificPodData.ObjectMeta.Labels,
				Annotations:       getSpecificPodData.ObjectMeta.Annotations,
				CreationTimestamp: getSpecificPodData.ObjectMeta.CreationTimestamp,
				DeletionTimestamp: getSpecificPodData.ObjectMeta.DeletionTimestamp,
				OwnerReferences:   getSpecificPodData.ObjectMeta.OwnerReferences,
			},
		}

		customGetSpecificPodDataJson, _ := json.Marshal(customGetSpecificPodData)
		ctx.JSON(http.StatusOK, string(customGetSpecificPodDataJson))
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
