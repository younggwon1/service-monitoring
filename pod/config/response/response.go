package response

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ResponseAllPodData struct {
	Name           string `json:"name"`
	NameSpace      string `json:"namespace"`
	Status         string `json:"status"`
	Age            string `json:"age"`
	Restarts       int32  `json:"restarts"`
	Ready          int    `json:"ready"`
	TotalContainer int    `json:"totalcontainer"`
}

type ResponseSpecificPodData struct {
	Name      string                        `json:"name"`
	Namespace string                        `json:"namespace"`
	Status    ResponseSpecificPodStatusData `json:"status"`
	Spec      ResponseSpecificPodSpecData   `json:"spec"`
	Meta      ResponseSpecificPodMetaData   `json:"meta"`
}

type ResponseSpecificPodStatusData struct {
	InitContainerStatuses []corev1.ContainerStatus `json:"initContainerStatuses"`
	ContainerStatuses     []corev1.ContainerStatus `json:"containerStatuses"`
	HostIP                string                   `json:"hostIP"`
	PodIP                 string                   `json:"podIP"`
	Phase                 corev1.PodPhase          `json:"phase"`
	QOSClass              corev1.PodQOSClass       `json:"qosClass"`
	Message               string                   `json:"message"`
}

type ResponseSpecificPodSpecData struct {
	InitContainerStatuses []corev1.Container         `json:"initContainerStatuses"`
	ContainerStatuses     []corev1.Container         `json:"containerStatuses"`
	RestartPolicy         corev1.RestartPolicy       `json:"restartPolicy"`
	SchedulerName         string                     `json:"schedulerName"`
	SecurityContext       *corev1.PodSecurityContext `json:"securityContext"`
	ServiceAccountName    string                     `json:"serviceAccountName"`
	Volumes               []corev1.Volume            `json:"volumes"`
}

type ResponseSpecificPodMetaData struct {
	Name              string                  `json:"name"`
	Namespace         string                  `json:"namespace"`
	Labels            map[string]string       `json:"labels"`
	Annotations       map[string]string       `json:"annotations"`
	CreationTimestamp metav1.Time             `json:"creationTimestamp"`
	DeletionTimestamp *metav1.Time            `json:"deletionTimestamp"`
	OwnerReferences   []metav1.OwnerReference `json:"ownerReferences"`
}
