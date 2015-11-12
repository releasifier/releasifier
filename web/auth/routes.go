package auth

import "github.com/pressly/chi"

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", login)
	r.Post("/permission", setPermission)
	r.Get("/logout", logout)

	return r
}
