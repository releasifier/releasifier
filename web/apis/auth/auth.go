package auth

import (
	"fmt"
	"net/http"

	"github.com/alinz/releasifier/data"
	"github.com/alinz/releasifier/lib/utils"
	"github.com/alinz/releasifier/web/constants"
	"github.com/alinz/releasifier/web/security"
	"golang.org/x/net/context"
)

func login(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	loginReq := ctx.Value(constants.CtxKeyParsedBody).(*loginRequest)
	user, err := data.DB.User.FindByEmailPassword(*loginReq.Email, *loginReq.Password)

	if err != nil {
		utils.Respond(w, 401, fmt.Errorf("unauthorized"))
		return
	}

	claims := map[string]interface{}{"user_id": fmt.Sprintf("%v", user.ID)}
	_, tokenStr, err := security.TokenAuth.Encode(claims)
	if err != nil {
		security.RemoveJwtCookie(w)
		utils.Respond(w, 401, fmt.Errorf("Authorization failed."))
		return
	}

	security.SetJwtCookie(tokenStr, w)
	utils.Respond(w, 200, user)
}

func logout(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	security.RemoveJwtCookie(w)
	utils.Respond(w, 200, "")
}
