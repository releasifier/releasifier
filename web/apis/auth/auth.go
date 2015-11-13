package auth

import (
	"net/http"

	"github.com/alinz/releasifier/data"
	"github.com/alinz/releasifier/lib/logme"
	"github.com/alinz/releasifier/lib/utils"
	m "github.com/alinz/releasifier/web/middlewares"
	"golang.org/x/net/context"
)

func login(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	loginReq := ctx.Value(m.BodyParserCtxKey).(*loginRequest)
	user, err := data.DB.User.FindByEmailPassword(loginReq.Email, loginReq.Password)

	logme.Info(loginReq)

	if err != nil {
		utils.Respond(w, 401, err)
		return
	}

	utils.Respond(w, 200, user)
}

func logout(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, 200, "bye")
}

func setPermissions(ctx context.Context, w http.ResponseWriter, r *http.Request) {

}

func getPermissions(ctx context.Context, w http.ResponseWriter, r *http.Request) {

}
