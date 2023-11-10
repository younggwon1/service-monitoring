package request

type RequestSpecificPodData struct {
	Name      string `form:"name"`
	NameSpace string `form:"namespace"`
}

type RequestSpecificPodLogsData struct {
	Name      string `form:"name"`
	NameSpace string `form:"namespace"`
}
