package data

import (
	"path/filepath"
	"time"

	"github.com/alinz/releasifier/common"
	internalErrors "github.com/alinz/releasifier/errors"
	"github.com/alinz/releasifier/logme"
	"upper.io/bond"
)

//BundleStore store for bundle
type BundleStore struct {
	bond.Store
}

//UploadBundles try to save them in to db
func (s BundleStore) UploadBundles(releaseID, appID, userID int64, fileInfos []*common.FileInfo) ([]*Bundle, error) {
	tx, err := DB.NewTransaction()
	if err != nil {
		return nil, internalErrors.ErrorSomethingWentWrong
	}
	defer tx.Close()

	b := tx.Builder()
	q := b.
		Select("releases.id", "releases.app_id", "releases.platform", "releases.note", "releases.version", "releases.created_at", "releases.private").
		From("releases").
		Join("apps").
		On("apps.id=releases.app_id").
		Join("apps_users_permissions").
		On("apps.id=apps_users_permissions.app_id").
		Where("releases.id=? AND apps_users_permissions.user_id=? AND apps_users_permissions.app_id=? AND apps_users_permissions.permission!=?", releaseID, userID, appID, MEMBER)

	var release *Release

	err = q.Iterator().One(&release)

	if err != nil {
		logme.Warn(err.Error())
		return nil, internalErrors.ErrorReleaseNotFound
	}

	if !release.Private {
		return nil, internalErrors.ErrorReleaseLocked
	}

	var bundles []*Bundle

	for _, fileInfo := range fileInfos {
		var fileType FileType

		extension := filepath.Ext(fileInfo.Filename)
		if extension == "jsbundle" {
			fileType = CODE
		} else {
			fileType = IMAGE
		}

		bundle := &Bundle{
			ReleaseID: releaseID,
			Hash:      fileInfo.Hash,
			Name:      fileInfo.Filename,
			Type:      fileType,
			CreatedAt: time.Now().UTC().Truncate(time.Second),
		}

		err = tx.Save(bundle)
		if err != nil {
			logme.Warn(err.Error())
			return nil, internalErrors.ErrorSomethingWentWrong
		}

		bundles = append(bundles, bundle)
	}

	if err = tx.Commit(); err != nil {
		logme.Warn(err.Error())
		return nil, internalErrors.ErrorDuplicateName
	}

	return bundles, nil
}
