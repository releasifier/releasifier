package apps

import (
	"net/http"
	"strconv"

	"github.com/alinz/releasifier/data"
	"github.com/alinz/releasifier/lib/utils"
	"github.com/alinz/releasifier/web/constants"
	"github.com/dgrijalva/jwt-go"
	"github.com/pressly/chi"
	"golang.org/x/net/context"
)

func getAllApps(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	token := ctx.Value(constants.CtxKeyJwtToken).(*jwt.Token)
	userIDStr := token.Claims["user_id"].(string)
	userID, _ := strconv.ParseInt(userIDStr, 10, 64)

	apps, err := data.DB.App.FindAllByUserID(userID)

	if err != nil {
		utils.Respond(w, 400, err)
	} else {
		utils.Respond(w, 200, apps)
	}
}

func createApp(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	createAppReq := ctx.Value(constants.CtxKeyParsedBody).(*createAppRequest)

	app, err := data.DB.App.CreateNewApp(1, createAppReq.Name, "", "")

	if err == nil {
		utils.Respond(w, 200, app)
	} else {
		utils.Respond(w, 400, err)
	}
}

func updateApp(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	appID := chi.URLParams(ctx)["appID"]
	utils.Respond(w, 200, "update app with id of "+appID)
}

func removeApp(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	appID := chi.URLParams(ctx)["appID"]
	utils.Respond(w, 200, "delete an app with id of "+appID)
}
