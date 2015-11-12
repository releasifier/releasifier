package auth

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func loginRequestBuilder() interface{} {
	return &loginRequest{}
}
