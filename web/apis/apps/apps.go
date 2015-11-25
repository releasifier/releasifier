package apps

import (
	"fmt"
	"net/http"

	"github.com/alinz/releasifier/data"
	"github.com/alinz/releasifier/lib/utils"
	"github.com/alinz/releasifier/web/constants"
	"github.com/alinz/releasifier/web/util"
	"github.com/pressly/chi"
	"golang.org/x/net/context"
)

func getAllApps(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	userID, _ := util.GetUserIDFromContext(ctx)

	apps, err := data.DB.App.FindAllApps(userID)

	if err != nil {
		utils.Respond(w, 400, err)
	} else {
		utils.Respond(w, 200, apps)
	}
}

func getApp(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	userID, _ := util.GetUserIDFromContext(ctx)
	appID, _ := util.GetParamValueAsID(ctx, "appID")

	utils.Respond(w, 200, fmt.Sprintf("get app with id of %d for user %d", appID, userID))
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
