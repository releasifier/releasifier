package bundles

import (
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/alinz/releasifier/config"
	"github.com/alinz/releasifier/data"
	"github.com/alinz/releasifier/lib/crypto"
	"github.com/alinz/releasifier/lib/utils"
	"github.com/alinz/releasifier/logme"
	"github.com/alinz/releasifier/web/util"
	"github.com/pressly/chi"

	"golang.org/x/net/context"
)

func copyDataToFile(in io.Reader, to string) error {
	out, err := os.Create(to)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, in)

	if err != nil {
		return err
	}

	return nil
}

func uploadBundle(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	//userID, _ := util.GetUserIDFromContext(ctx)
	//appID, _ := util.GetParamValueAsID(ctx, "appID")

	if err := r.ParseMultipartForm(config.Conf.FileUpload.MaxSize); err != nil {
		utils.Respond(w, http.StatusForbidden, err)
		return
	}

	type FileInfo struct {
		Filename string
		Src      string
	}

	fileInfos := make([]*FileInfo, 0)

	var path string

	//saving all the files into temp folder
	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			path = config.Conf.FileUpload.Temp + fileHeader.Filename
			if err := copyDataToFile(file, path); err != nil {
				utils.Respond(w, http.StatusExpectationFailed, err)
				return
			}

			file.Close()

			fileInfos = append(fileInfos, &FileInfo{
				Filename: fileHeader.Filename,
				Src:      path,
			})
		}
	}

	for _, fileInfo := range fileInfos {
		hash, err := crypto.HashFile(fileInfo.Src)
		if err != nil {
			utils.Respond(w, http.StatusExpectationFailed, err)
			return
		}
		logme.Info(hash)
	}

	utils.Respond(w, http.StatusOK, nil)
}

func getBundleByHash(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	userID, _ := util.GetUserIDFromContext(ctx)
	appID, _ := util.GetParamValueAsID(ctx, "appID")
	hash := chi.URLParams(ctx)["hash"]

	//check if current logged in user has admin or owner permission
	if data.DB.App.HasPermission(userID, appID, data.OWNER, data.ADMIN) {
		utils.Respond(w, http.StatusForbidden, errors.New("you don't have write access to this app"))
	}

	bundle, err := data.DB.Bundle.FindByHash(hash)

	if bundle == nil || err != nil {
		utils.Respond(w, http.StatusNotFound, errors.New("bundle with requested hash not found"))
	}

	utils.Respond(w, http.StatusOK, bundle)
}
