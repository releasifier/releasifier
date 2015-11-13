package security

import (
	"net/http"

	"github.com/goware/jwtauth"
)

var TokenAuth *jwtauth.JwtAuth

//SetJwtCookie is going to set during security.Setup.
//in this way, this function can access config
var SetJwtCookie func(token string, w http.ResponseWriter)

func setupTokenAuth(secretKey string) {
	TokenAuth = jwtauth.New("HS256", []byte(secretKey), nil)
}
