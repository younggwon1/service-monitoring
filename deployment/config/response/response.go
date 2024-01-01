package response

type ResponseErrorDeployments struct {
	Name      string `json:"name"`
	NameSpace string `json:"namespace"`
	Reason    string `json:"reason"`
	Message   string `json:"message"`
	Age       string `json:"age"`
}

type ResponseDeleteErrorDeployments struct {
	Name      string `json:"name"`
	NameSpace string `json:"namespace"`
	Reason    string `json:"reason"`
	Message   string `json:"message"`
	Age       string `json:"age"`
}
