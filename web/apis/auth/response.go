package auth

type loginResponse struct {
	ID  int64  `json:"id"`
	Jwt string `json:"jwt"`
}
