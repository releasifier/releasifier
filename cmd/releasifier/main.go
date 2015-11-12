package main

import (
	"net/http"

	"github.com/pressly/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	http.ListenAndServe(":7331", r)
}
