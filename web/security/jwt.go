package security

import (
	"net/http"

	"github.com/goware/jwtauth"
)

//TokenAuth is uses to generate jwt token
var TokenAuth *jwtauth.JwtAuth

//SetJwtCookie is going to set during security.Setup.
//in this way, this function can access config
var SetJwtCookie func(token string, w http.ResponseWriter)

//RemoveJwtCookie is going to set during security.Setup.
//in this way, this function can access config
var RemoveJwtCookie func(w http.ResponseWriter)

func setupTokenAuth(secretKey string) {
	TokenAuth = jwtauth.New("HS256", []byte(secretKey), nil)
}
