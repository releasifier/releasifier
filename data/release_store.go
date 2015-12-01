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

//FindAllReleases returns all release for specific user and app. Only memebers of that app can see the list
func (s ReleaseStore) FindAllReleases(appID, userID int64) ([]*Release, error) {
	b := s.Session().Builder()
	q := b.
		Select("releases.id", "releases.platform", "releases.note", "releases.version", "releases.created_at", "releases.private").
		From("releases").
		Join("apps").
		On("apps.id=releases.app_id").
		Join("apps_users_permissions").
		On("apps.id=apps_users_permissions.app_id").
		Where("apps_users_permissions.user_id=? AND apps_users_permissions.app_id=?", userID, appID)

	var releases []*Release

	err := q.Iterator().All(&releases)

	if err != nil {
		return nil, err
	}

	if releases == nil {
		releases = make([]*Release, 0)
	}

	return releases, nil
}
