package auth

import (
	"fmt"
	"net/http"

	"github.com/alinz/releasifier/data"
	"github.com/alinz/releasifier/errors"
	"github.com/alinz/releasifier/lib/utils"
	"github.com/alinz/releasifier/web/constants"
	"github.com/alinz/releasifier/web/security"
	"golang.org/x/net/context"
)

func loginUser(w http.ResponseWriter, email, password string) {
	user, err := data.DB.User.FindByEmailPassword(email, password)

	if err != nil {
		utils.RespondEx(w, nil, 0, errors.ErrorAuthorizeAccess)
		return
	}

	claims := map[string]interface{}{"user_id": fmt.Sprintf("%v", user.ID)}
	_, tokenStr, err := security.TokenAuth.Encode(claims)
	if err != nil {
		security.RemoveJwtCookie(w)
		utils.RespondEx(w, nil, 0, errors.ErrorAuthorizeAccess)
		return
	}

	security.SetJwtCookie(tokenStr, w)
	utils.RespondEx(w, loginResponse{ID: user.ID, Jwt: tokenStr}, 0, nil)
}

func register(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	registerReq := ctx.Value(constants.CtxKeyParsedBody).(*registerRequest)

	_, err := data.DB.User.Create(*registerReq.Fullname, *registerReq.Email, *registerReq.Password)
	if err != nil {
		utils.RespondEx(w, nil, 0, err)
		return
	}

	loginUser(w, *registerReq.Email, *registerReq.Password)
}

func login(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	loginReq := ctx.Value(constants.CtxKeyParsedBody).(*loginRequest)
	loginUser(w, *loginReq.Email, *loginReq.Password)
}

func logout(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	security.RemoveJwtCookie(w)
	utils.Respond(w, 200, nil)
}
