package apps

import (
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/alinz/releasifier/common"
	"github.com/alinz/releasifier/data"
	internalErrors "github.com/alinz/releasifier/errors"
	"github.com/alinz/releasifier/logme"
	"github.com/alinz/releasifier/web/util"

	"github.com/alinz/releasifier/config"
	"github.com/alinz/releasifier/lib/crypto"
	"github.com/alinz/releasifier/lib/utils"

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

func uuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func uploadBundles(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	userID, _ := util.GetUserIDFromContext(ctx)
	appID, _ := util.GetParamValueAsID(ctx, "appID")
	releaseID, _ := util.GetParamValueAsID(ctx, "releaseID")

	//check if user has permission admin or owner
	if !data.DB.App.HasPermission(appID, userID, data.ADMIN, data.OWNER) {
		utils.RespondEx(w, nil, 0, internalErrors.ErrorAuthorizeAccess)
		return
	}

	if err := r.ParseMultipartForm(config.Conf.FileUpload.MaxSize); err != nil {
		utils.RespondEx(w, nil, 0, internalErrors.ErrorSomethingWentWrong)
		logme.Warn(err.Error())
		return
	}

	var fileInfos []*common.FileInfo
	var path string

	fileInfos = make([]*common.FileInfo, 0)

	//saving all the files into temp folder
	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			path = config.Conf.FileUpload.Temp + uuid()
			if err := copyDataToFile(file, path); err != nil {
				utils.RespondEx(w, nil, 0, internalErrors.ErrorSomethingWentWrong)
				logme.Warn(err.Error())
				return
			}

			file.Close()

			hash, err := crypto.HashFile(path)
			if err != nil {
				utils.RespondEx(w, nil, 0, internalErrors.ErrorSomethingWentWrong)
				logme.Warn(err.Error())
				return
			}

			fileInfos = append(fileInfos, &common.FileInfo{
				Filename:     fileHeader.Filename,
				Hash:         hash,
				TempLocation: path,
			})
		}
	}

	bundles, err := data.DB.Bundle.UploadBundles(releaseID, appID, userID, fileInfos)

	if err != nil {
		logme.Warn(err.Error())
		//remove all temp files
		for _, fileInfo := range fileInfos {
			os.Remove(fileInfo.TempLocation)
		}
	} else {
		for _, fileInfo := range fileInfos {
			err = os.Rename(fileInfo.TempLocation, config.Conf.FileUpload.Bundle+fileInfo.Hash)
			logme.Info(fileInfo.TempLocation)
			if err != nil {
				logme.Warn(err.Error())
			}
		}
	}

	utils.RespondEx(w, bundles, 0, err)
}

func getAllBundles(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	userID, _ := util.GetUserIDFromContext(ctx)
	appID, _ := util.GetParamValueAsID(ctx, "appID")
	releaseID, _ := util.GetParamValueAsID(ctx, "releaseID")

	bundles, err := data.DB.Bundle.FindAllBundles(releaseID, appID, userID)
	utils.RespondEx(w, bundles, 0, err)
}

func deleteBundle(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	userID, _ := util.GetUserIDFromContext(ctx)
	appID, _ := util.GetParamValueAsID(ctx, "appID")
	releaseID, _ := util.GetParamValueAsID(ctx, "releaseID")
	bundleID, _ := util.GetParamValueAsID(ctx, "bundleID")

	err := data.DB.Bundle.RemoveBundle(bundleID, releaseID, appID, userID)
	utils.RespondEx(w, nil, 0, err)
}
