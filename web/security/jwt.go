package security

import "net/http"

//TokenAuth is uses to generate jwt token
var TokenAuth *JwtAuth

//SetJwtCookie is going to set during security.Setup.
//in this way, this function can access config
var SetJwtCookie func(token string, w http.ResponseWriter)

//RemoveJwtCookie is going to set during security.Setup.
//in this way, this function can access config
var RemoveJwtCookie func(w http.ResponseWriter)

func setupTokenAuth(secretKey string) {
	TokenAuth = New("HS256", []byte(secretKey), nil)
}
