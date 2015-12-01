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

	r.Route("/:appID", func(r chi.Router) {
		r.Get("/", getApp)
		r.Patch("/", m.BodyParser(updateAppRequestBuilder, 2^14+100), updateApp)
		r.Delete("/", removeApp)

		r.Route("/token", func(r chi.Router) {
			r.Post("/", m.BodyParser(generateAppTokenRequestBuilder, 100), generateAppToken)
			r.Put("/", m.BodyParser(appTokenRequestBuilder, 200), acceptAppToken)
		})

		r.Route("/releases", func(r chi.Router) {
			r.Post("/", m.BodyParser(createReleaseRequestBuilder, 1024), createRelease)
		})
	})

	return r
}
