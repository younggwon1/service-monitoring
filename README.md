# Service Monitoring

## API 명세서
### 1. Get API
1. 전체 POD Data 를 조회하는 API
Request
- Request Syntax
http://{SERVER_URL}/v1/podData/pods
- Request Header
- Request Elements

Response
- Response Syntax
- Response Header
- Response Elements
2. 특정 POD Data 를 조회하는 API
Request
- Request Syntax
- Request Header
- Request Elements

Response
- Response Syntax
- Response Header
- Response Elements
3. 특정 POD Log 를 조회하는 API
Request
- Request Syntax
- Request Header
- Request Elements

Response
- Response Syntax
- Response Header
- Response Elements
### 2. Posts API


해야할 일
interface
- https://www.jetbrains.com/go/guide/tutorials/rest_api_series/gin/
 - pod, deployment, cronjob 등의 interface 를 만들면 좋지 않을까??
swagger
- https://github.com/swaggo/swag




		// getSpecificPodData.Status
		// fmt.Println("getSpecificPodData.Status.ContainerStatuses : ", getSpecificPodData.Status.ContainerStatuses)
		// fmt.Println("getSpecificPodData.Status.ContainerStatuses : ", reflect.TypeOf(getSpecificPodData.Status.ContainerStatuses))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.Status.HostIP : ", getSpecificPodData.Status.HostIP)
		// fmt.Println("getSpecificPodData.Status.HostIP : ", reflect.TypeOf(getSpecificPodData.Status.HostIP))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.Status.InitContainerStatuses : ", getSpecificPodData.Status.InitContainerStatuses)
		// fmt.Println("getSpecificPodData.Status.InitContainerStatuses : ", reflect.TypeOf(getSpecificPodData.Status.InitContainerStatuses))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.Status.Message : ", getSpecificPodData.Status.Message)
		// fmt.Println("getSpecificPodData.Status.Message : ", reflect.TypeOf(getSpecificPodData.Status.Message))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.Status.Phase : ", getSpecificPodData.Status.Phase)
		// fmt.Println("getSpecificPodData.Status.Phase : ", reflect.TypeOf(getSpecificPodData.Status.Phase))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.Status.PodIP : ", getSpecificPodData.Status.PodIP)
		// fmt.Println("getSpecificPodData.Status.PodIP : ", reflect.TypeOf(getSpecificPodData.Status.PodIP))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.Status.QOSClass : ", getSpecificPodData.Status.QOSClass)
		// fmt.Println("getSpecificPodData.Status.QOSClass : ", reflect.TypeOf(getSpecificPodData.Status.QOSClass))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.Status.StartTime : ", getSpecificPodData.Status.StartTime)
		// fmt.Println("getSpecificPodData.Status.StartTime : ", reflect.TypeOf(getSpecificPodData.Status.StartTime))

		// getSpecificPodData.Spec
		// fmt.Println("getSpecificPodData.Spec.InitContainers : ", getSpecificPodData.Spec.InitContainers)
		// fmt.Println("getSpecificPodData.Spec.InitContainers : ", reflect.TypeOf(getSpecificPodData.Spec.InitContainers))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.Spec.Containers : ", getSpecificPodData.Spec.Containers)
		// fmt.Println("getSpecificPodData.Spec.Containers : ", reflect.TypeOf(getSpecificPodData.Spec.Containers))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.Spec.RestartPolicy : ", getSpecificPodData.Spec.RestartPolicy)
		// fmt.Println("getSpecificPodData.Spec.RestartPolicy : ", reflect.TypeOf(getSpecificPodData.Spec.RestartPolicy))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.Spec.SchedulerName : ", getSpecificPodData.Spec.SchedulerName)
		// fmt.Println("getSpecificPodData.Spec.SchedulerName : ", reflect.TypeOf(getSpecificPodData.Spec.SchedulerName))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.Spec.SecurityContext : ", getSpecificPodData.Spec.SecurityContext)
		// fmt.Println("getSpecificPodData.Spec.SecurityContext : ", reflect.TypeOf(getSpecificPodData.Spec.SecurityContext))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.Spec.ServiceAccountName : ", getSpecificPodData.Spec.ServiceAccountName)
		// fmt.Println("getSpecificPodData.Spec.ServiceAccountName : ", reflect.TypeOf(getSpecificPodData.Spec.ServiceAccountName))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.Spec.Volumes : ", getSpecificPodData.Spec.Volumes)
		// fmt.Println("getSpecificPodData.Spec.Volumes : ", reflect.TypeOf(getSpecificPodData.Spec.Volumes))

		// getSpecificPodData.ObjectMeta
		// fmt.Println("getSpecificPodData.ObjectMeta.Name : ", getSpecificPodData.ObjectMeta.Name)
		// fmt.Println("getSpecificPodData.Spec.ServiceAccountName : ", reflect.TypeOf(getSpecificPodData.ObjectMeta.Name))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.ObjectMeta.Namespace : ", getSpecificPodData.ObjectMeta.Namespace)
		// fmt.Println("getSpecificPodData.Spec.ServiceAccountName : ", reflect.TypeOf(getSpecificPodData.ObjectMeta.Namespace))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.ObjectMeta.Labels : ", getSpecificPodData.ObjectMeta.Labels)
		// fmt.Println("getSpecificPodData.Spec.ServiceAccountName : ", reflect.TypeOf(getSpecificPodData.ObjectMeta.Labels))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.ObjectMeta.Annotations : ", getSpecificPodData.ObjectMeta.Annotations)
		// fmt.Println("getSpecificPodData.Spec.ServiceAccountName : ", reflect.TypeOf(getSpecificPodData.ObjectMeta.Annotations))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.ObjectMeta.CreationTimestamp : ", getSpecificPodData.ObjectMeta.CreationTimestamp)
		// fmt.Println("getSpecificPodData.Spec.ServiceAccountName : ", reflect.TypeOf(getSpecificPodData.ObjectMeta.CreationTimestamp))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.ObjectMeta.DeletionTimestamp : ", getSpecificPodData.ObjectMeta.DeletionTimestamp)
		// fmt.Println("getSpecificPodData.Spec.ServiceAccountName : ", reflect.TypeOf(getSpecificPodData.ObjectMeta.DeletionTimestamp))
		// fmt.Println("------------------------------------------------------------------")
		// fmt.Println("getSpecificPodData.ObjectMeta.OwnerReferences : ", getSpecificPodData.ObjectMeta.OwnerReferences)
		// fmt.Println("getSpecificPodData.Spec.ServiceAccountName : ", reflect.TypeOf(getSpecificPodData.ObjectMeta.OwnerReferences))
		// fmt.Println("------------------------------------------------------------------")