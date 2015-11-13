package app

import (
	"github.com/alinz/releasifier/web/security"
	"github.com/pressly/chi"
)

//Routes returns chi's Router for App APIs
func Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(security.TokenAuth.Handle("state"))

	r.Get("/", getAllApps)
	r.Post("/", createApp)
	r.Put("/:appID", updateApp)
	r.Delete("/:appID", removeApp)

	return r
}
