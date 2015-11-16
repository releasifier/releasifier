package apps

type createAppRequest struct {
	Name string `json:"name"`
}

func createAppRequestBuilder() interface{} {
	return &createAppRequest{}
}
