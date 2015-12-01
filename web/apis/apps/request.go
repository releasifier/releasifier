package apps

import "github.com/alinz/releasifier/data"

type createAppRequest struct {
	Name *string `json:"name,required"`
}

func createAppRequestBuilder() interface{} {
	return &createAppRequest{}
}

type updateAppRequest struct {
	Name       *string `json:"name"`
	PublicKey  *string `json:"public_key"`
	PrivateKey *string `json:"private_key"`
	Private    *bool   `json:"private"`
}

func updateAppRequestBuilder() interface{} {
	return &updateAppRequest{}
}

type generateAppTokenRequest struct {
	Permission *data.Permission `json:"permission,required"`
}

func generateAppTokenRequestBuilder() interface{} {
	return &generateAppTokenRequest{}
}

type appTokenRequest struct {
	Token *string `json:"token,required"`
}

func appTokenRequestBuilder() interface{} {
	return &appTokenRequest{}
}
