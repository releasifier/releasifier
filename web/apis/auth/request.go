package auth

type loginRequest struct {
	Email    *string `json:"email,required"`
	Password *string `json:"password,required"`
}

func loginRequestBuilder() interface{} {
	return &loginRequest{}
}
