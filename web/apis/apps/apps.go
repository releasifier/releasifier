package apps

import (
	"net/http"

	"github.com/alinz/releasifier/data"
	"github.com/alinz/releasifier/lib/utils"
	"github.com/alinz/releasifier/web/constants"
	"github.com/pressly/chi"
	"golang.org/x/net/context"
)

func getAllApps(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	//token := ctx.Value(constants.CtxKeyJwtToken).(*jwt.Token)
	//userID := token.Claims["user_id"].(string)

	utils.Respond(w, 200, "all apps")
}

func createApp(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	createAppReq := ctx.Value(constants.CtxKeyParsedBody).(*createAppRequest)

	app, err := data.DB.App.CreateNewApp(createAppReq.Name)

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
