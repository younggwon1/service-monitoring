package response

type ResponseAllDeployments struct {
	Name        string `json:"name"`
	NameSpace   string `json:"namespace"`
	Age         string `json:"age"`
	Current     int32  `json:"current"`
	Desired     int32  `json:"desired"`
	UptoDate    int32  `json:"uptodate"`
	Available   int32  `json:"available"`
	UnAvailable int32  `json:"unavailable"`
}
