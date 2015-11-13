package app

import (
	"github.com/alinz/releasifier/web/security"
	"github.com/pressly/chi"
)

//Routes returns chi's Router for Auth APIs
func Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(security.TokenAuth.Handle("state"))

	r.Get("/test", test)

	return r
}
