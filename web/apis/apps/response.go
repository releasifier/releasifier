package apps

type generateAppTokenResponse struct {
	Token string `json:"token"`
}

func generateAppTokenResponseBuilder() interface{} {
	return &generateAppTokenResponse{}
}
