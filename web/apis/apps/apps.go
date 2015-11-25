package apps

import (
	"net/http"

	"github.com/alinz/releasifier/data"
	"github.com/alinz/releasifier/lib/utils"
	"github.com/alinz/releasifier/web/constants"
	"github.com/alinz/releasifier/web/util"
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

	app, err := data.DB.App.FindApp(appID, userID)

	if err == nil {
		utils.Respond(w, 200, app)
	} else {
		utils.Respond(w, 400, err)
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
	userID, _ := util.GetUserIDFromContext(ctx)
	appID, _ := util.GetParamValueAsID(ctx, "appID")
	updateAppReq := ctx.Value(constants.CtxKeyParsedBody).(*updateAppRequest)

	err := data.DB.App.UpdateApp(appID, updateAppReq.Name, updateAppReq.PublicKey, updateAppReq.PrivateKey, userID)

	if err == nil {
		utils.Respond(w, 200, nil)
	} else {
		utils.Respond(w, 400, err)
	}
}

func removeApp(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	userID, _ := util.GetUserIDFromContext(ctx)
	appID, _ := util.GetParamValueAsID(ctx, "appID")

	err := data.DB.App.RemoveApp(appID, userID)
	if err == nil {
		utils.Respond(w, 200, nil)
	} else {
		utils.Respond(w, 400, err)
	}
}
