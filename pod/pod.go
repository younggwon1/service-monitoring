package pod

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"

	request "github.com/younggwon1/service-monitoring/pod/config/request"
	response "github.com/younggwon1/service-monitoring/pod/config/response"
)

func AllPods(ctx *gin.Context, clientSet *kubernetes.Clientset) {
	getAllPodsData, err := clientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var getAllPodsDataList []response.ResponseAllPodData
	for _, pod := range getAllPodsData.Items {
		podInfo := &response.ResponseAllPodData{
			Name:      pod.Name,
			NameSpace: pod.Namespace,
			Status:    string(pod.Status.Phase),
			Age:       pod.Status.StartTime.Format("2006-01-02 15:04:05"),
		}
		for _, containerStatus := range pod.Status.ContainerStatuses {
			if containerStatus.State.Running != nil && containerStatus.Ready {
				podInfo.Restarts = containerStatus.RestartCount
				podInfo.Ready++
				podInfo.TotalContainer = len(pod.Status.ContainerStatuses)
			}
		}
		getAllPodsDataList = append(getAllPodsDataList, *podInfo)
	}

	customGetPodsDataJson, _ := json.Marshal(getAllPodsDataList)
	ctx.JSON(http.StatusOK, string(customGetPodsDataJson))
}

func SpecificPod(ctx *gin.Context, clientSet *kubernetes.Clientset) {
	var RequestSpecificPodData request.RequestSpecificPodData

	if err := ctx.ShouldBindQuery(&RequestSpecificPodData); err == nil {
		getSpecificPodData, err := clientSet.CoreV1().Pods(RequestSpecificPodData.NameSpace).Get(context.TODO(), RequestSpecificPodData.Name, metav1.GetOptions{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		customGetSpecificPodData := &response.ResponseSpecificPodData{
			Name:      getSpecificPodData.Name,
			Namespace: getSpecificPodData.Namespace,
			Status: response.ResponseSpecificPodStatusData{
				InitContainerStatuses: getSpecificPodData.Status.InitContainerStatuses,
				ContainerStatuses:     getSpecificPodData.Status.InitContainerStatuses,
				HostIP:                getSpecificPodData.Status.HostIP,
				PodIP:                 getSpecificPodData.Status.PodIP,
				Phase:                 getSpecificPodData.Status.Phase,
				QOSClass:              getSpecificPodData.Status.QOSClass,
				Message:               getSpecificPodData.Status.Message,
			},
			Spec: response.ResponseSpecificPodSpecData{
				InitContainerStatuses: getSpecificPodData.Spec.InitContainers,
				ContainerStatuses:     getSpecificPodData.Spec.Containers,
				RestartPolicy:         getSpecificPodData.Spec.RestartPolicy,
				SchedulerName:         getSpecificPodData.Spec.SchedulerName,
				SecurityContext:       getSpecificPodData.Spec.SecurityContext,
				ServiceAccountName:    getSpecificPodData.Spec.ServiceAccountName,
				Volumes:               getSpecificPodData.Spec.Volumes,
			},
			Meta: response.ResponseSpecificPodMetaData{
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func SpecificPodLogs(ctx *gin.Context, clientSet *kubernetes.Clientset) {
	var RequestSpecificPodLogsData request.RequestSpecificPodLogsData
	podLogOpts := corev1.PodLogOptions{Follow: true}

	if err := ctx.ShouldBindQuery(&RequestSpecificPodLogsData); err == nil {
		req := clientSet.CoreV1().Pods(RequestSpecificPodLogsData.NameSpace).GetLogs(RequestSpecificPodLogsData.Name, &podLogOpts)
		podLogs, err := req.Stream(context.TODO())
		if err != nil {
			panic(err.Error())
		}

		defer podLogs.Close()

		buffer := bufio.NewReader(podLogs)
		for {
			val, readErr := buffer.ReadString('\n')
			if readErr == io.EOF {
				fmt.Println("EOF")
				break
			}
			fmt.Println(val)
		}
	}
}

func LiveStreamSpecificPodResourceUsage(ctx *gin.Context, metrics *metrics.Clientset) {
	var RequestSpecificPodData request.RequestSpecificPodData
	// Use a time-based ticker to periodically fetch pod metrics.
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// Fetch the pod metrics.
			if err := ctx.ShouldBindQuery(&RequestSpecificPodData); err == nil {
				podMetrics, err := metrics.MetricsV1beta1().PodMetricses(RequestSpecificPodData.NameSpace).Get(context.TODO(), RequestSpecificPodData.Name, metav1.GetOptions{})
				if err != nil {
					fmt.Printf("Error fetching pod metrics: %v\n", err)
					continue
				}
				// Extract CPU and memory usage.
				for _, container := range podMetrics.Containers {
					fmt.Printf("Container: %s\n", container.Name)
					fmt.Printf("CPU Usage: %vm\n", podMetrics.Containers[0].Usage.Cpu().MilliValue())
					fmt.Printf("Memory Usage: %vm\n", podMetrics.Containers[0].Usage.Memory().MilliValue())
				}
			}
		}
	}
}
