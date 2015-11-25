package apps

type createAppRequest struct {
	Name string `json:"name"`
}

func createAppRequestBuilder() interface{} {
	return &createAppRequest{}
}

type updateAppRequest struct {
	Name       string `json:"name"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	Private    bool   `json:"private"`
}

func updateAppRequestBuilder() interface{} {
	return &updateAppRequest{}
}
