package auth

import (
	"net/http"

	"github.com/alinz/releasifier/lib/utils"
	"golang.org/x/net/context"
)

func login(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// t := ctx.Value("json_body")
	// body, ok := t.(*loginRequest)
	//
	//
}

func logout(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, 200, "bye")
}

func setPermissions(ctx context.Context, w http.ResponseWriter, r *http.Request) {

}

func getPermissions(ctx context.Context, w http.ResponseWriter, r *http.Request) {

}
