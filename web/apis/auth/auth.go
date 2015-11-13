package auth

import (
	"fmt"
	"net/http"

	"github.com/alinz/releasifier/data"
	"github.com/alinz/releasifier/lib/utils"
	m "github.com/alinz/releasifier/web/middlewares"
	"github.com/alinz/releasifier/web/security"
	"golang.org/x/net/context"
)

func login(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	loginReq := ctx.Value(m.BodyParserCtxKey).(*loginRequest)
	user, err := data.DB.User.FindByEmailPassword(loginReq.Email, loginReq.Password)

	if err != nil {
		utils.Respond(w, 401, fmt.Errorf("unauthorized"))
		return
	}

	security.SetJwtCookie("hello", w)

	utils.Respond(w, 200, user)
}

func logout(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, 200, "bye")
}
