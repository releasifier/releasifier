package releases

import (
	m "github.com/alinz/releasifier/web/middlewares"
	"github.com/alinz/releasifier/web/security"
	"github.com/pressly/chi"
)

//Routes returns chi's Router for Release's APIs
func Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(security.TokenAuth.Handle("state"))

	r.Post("/app/:appID", m.BodyParser(createReleaseRequestBuilder, 1024), createRelease)

	return r
}
