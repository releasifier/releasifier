package security

import (
	"net/http"

	"github.com/alinz/releasifier/config"
)

//Setup setups and configure security package.
func Setup(conf *config.Config) {
	setupTokenAuth(conf.JWT.SecretKey)

	SetJwtCookie = func(token string, w http.ResponseWriter) {
		cookie := http.Cookie{
			Name:     "jwt",
			Domain:   conf.JWT.Domain,
			Path:     conf.JWT.Path,
			Secure:   conf.JWT.Secure,
			MaxAge:   conf.JWT.MaxAge,
			HttpOnly: true,
			Value:    token,
		}
		http.SetCookie(w, &cookie)
	}

	RemoveJwtCookie = func(w http.ResponseWriter) {
		cookie := http.Cookie{
			Name:     "jwt",
			Domain:   conf.JWT.Domain,
			Path:     conf.JWT.Path,
			Secure:   conf.JWT.Secure,
			MaxAge:   -1,
			HttpOnly: true,
			Value:    "",
		}
		http.SetCookie(w, &cookie)
	}
}
