package auth

import "github.com/pressly/chi"

//Routes returns chi's Router for Auth APIs
func Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", login)
	r.Post("/permissions", setPermissions)
	r.Get("/permissions", getPermissions)
	r.Get("/logout", logout)

	return r
}
