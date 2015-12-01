package data

import (
	"time"

	internalErrors "github.com/alinz/releasifier/errors"
	"upper.io/bond"
)

//ReleaseStore store for release
type ReleaseStore struct {
	bond.Store
}

//CreateRelease create a release record. It always private.
func (s ReleaseStore) CreateRelease(version Version, platform Platform, note string, userID, appID int64) (*Release, error) {
	tx, err := DB.NewTransaction()
	if err != nil {
		return nil, err
	}
	defer tx.Close()

	if !DB.App.HasPermission(appID, userID, OWNER, ADMIN) {
		return nil, internalErrors.ErrorAuthorizeAccess
	}

	release := &Release{
		AppID:    appID,
		Platform: platform,
		Note:     note,
		Version:  version,
		CreateAt: time.Now().UTC().Truncate(time.Second),
		Private:  true,
	}

	if err = tx.Save(release); err != nil {
		return nil, internalErrors.ErrorDuplicateVersion
	}

	if err = tx.Commit(); err != nil {
		return nil, internalErrors.ErrorDuplicateVersion
	}

	return release, nil
}
