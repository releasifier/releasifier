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

	parseSecureIDParams := m.SecureIDParamsParser("appID")

	r.Get("/", getAllApps)
	r.Get("/:appID", parseSecureIDParams, getApp)
	r.Post("/", m.BodyParser(createAppRequestBuilder, 100), createApp)
	r.Put("/:appID", parseSecureIDParams, updateApp)
	r.Delete("/:appID", parseSecureIDParams, removeApp)

	return r
}
