package bundles

import (
	"github.com/alinz/releasifier/web/security"
	"github.com/pressly/chi"
)

//Routes returns chi's Router for Release's APIs
func Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(security.TokenAuth.Handle("state"))

	r.Post("/upload/app/:appID", uploadBundle)
	r.Get("/app/:appID/hash/:hash", getBundleByHash)

	return r
}
