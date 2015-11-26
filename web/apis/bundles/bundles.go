package bundles

import (
	"io"
	"net/http"
	"os"

	"github.com/alinz/releasifier/config"
	"github.com/alinz/releasifier/lib/utils"
	"github.com/alinz/releasifier/logme"

	"golang.org/x/net/context"
)

func readAndWrite(in io.Reader, fileName string) error {
	out, err := os.Create(config.Conf.FileUpload.Temp + fileName)
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

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			if err := readAndWrite(file, fileHeader.Filename); err != nil {
				logme.Fatal(err)
			}
		}
	}

	utils.Respond(w, http.StatusOK, nil)

	// fmt.Fprintf(w, "File uploaded successfully : ")
	// fmt.Fprintf(w, header.Filename)
}
