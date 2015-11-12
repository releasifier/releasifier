package web

import (
	"net/http"

	"github.com/alinz/releasifier/web/auth"
	"github.com/pressly/chi"
)

//New create all routers under one Handler
func New() http.Handler {
	r := chi.NewRouter()

	r.Mount("/auth", auth.Routes())

	return r
}
