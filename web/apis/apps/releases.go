package apps

import (
	"net/http"

	"github.com/alinz/releasifier/data"
	"github.com/alinz/releasifier/lib/utils"
	"github.com/alinz/releasifier/web/constants"
	"github.com/alinz/releasifier/web/util"
	"golang.org/x/net/context"
)

func getAllReleases(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	//get userID and appID
	userID, _ := util.GetUserIDFromContext(ctx)
	appID, _ := util.GetParamValueAsID(ctx, "appID")

	releases, err := data.DB.Release.FindAllReleases(userID, appID)

	utils.RespondEx(w, releases, 0, err)
}

func getRelease(ctx context.Context, w http.ResponseWriter, r *http.Request) {

}

func createRelease(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	//get userID and appID
	userID, _ := util.GetUserIDFromContext(ctx)
	appID, _ := util.GetParamValueAsID(ctx, "appID")

	//grabing release request
	createReleaseReq := ctx.Value(constants.CtxKeyParsedBody).(*createReleaseRequest)

	//try to create release and return created release record
	release, err := data.DB.Release.CreateRelease(*createReleaseReq.Version, *createReleaseReq.Platform, createReleaseReq.Note, userID, appID)
	if err == nil {
		utils.Respond(w, 200, release)
	} else {
		utils.Respond(w, 400, err)
	}
}
