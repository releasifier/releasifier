package apps

import (
	m "github.com/alinz/releasifier/web/middlewares"
	"github.com/alinz/releasifier/web/security"
	"github.com/pressly/chi"
)

//Routes returns chi's Router for App APIs
func Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(security.TokenAuth.Handle("state"))

	r.Post("/", m.BodyParser(createAppRequestBuilder, 100), createApp)
	r.Get("/", getAllApps)
	r.Get("/:appID", getApp)
	r.Patch("/:appID", m.BodyParser(updateAppRequestBuilder, 2^14+100), updateApp)
	r.Delete("/:appID", removeApp)

	r.Post("/:appID/token", m.BodyParser(generateAppTokenRequestBuilder, 100), generateAppToken)
	r.Put("/:appID/token", m.BodyParser(appTokenRequestBuilder, 200), acceptAppToken)

	return r
}
