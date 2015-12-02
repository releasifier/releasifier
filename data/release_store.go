package data

import (
	"fmt"
	"time"

	internalErrors "github.com/alinz/releasifier/errors"
	"github.com/alinz/releasifier/logme"
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

//UpdateRelease updates release record if it's private.
func (s ReleaseStore) UpdateRelease(note *string, platform *Platform, version *Version, releaseID, appID, userID int64) error {
	var requireToUpdate bool = true

	b := s.Session().Builder()
	q := b.
		Select("releases.id", "releases.app_id", "releases.platform", "releases.note", "releases.version", "releases.created_at", "releases.private").
		From("releases").
		Join("apps").
		On("apps.id=releases.app_id").
		Join("apps_users_permissions").
		On("apps.id=apps_users_permissions.app_id").
		Where("releases.id=? AND apps_users_permissions.user_id=? AND apps_users_permissions.app_id=? AND apps_users_permissions.permission!=?", releaseID, userID, appID, MEMBER)

	var release *Release

	err := q.Iterator().One(&release)

	if err != nil {
		return err
	}

	if release == nil {
		return internalErrors.ErrorReleaseNotFound
	}

	//once private becomes public, there is no turning back
	if !release.Private {
		return internalErrors.ErrorReleaseLocked
	}

	//check if one of `note`, `platform` or `version` has been requested to be updated
	if note != nil {
		release.Note = *note
		requireToUpdate = true
	}

	if platform != nil {
		release.Platform = *platform
		requireToUpdate = true
	}

	if version != nil {
		release.Version = *version
		requireToUpdate = true
	}

	//if required is set true it means that we can save the release.
	//it might adds a little bit performance gain! Still don't know!
	if requireToUpdate {
		if err = s.Save(release); err != nil {
			fmt.Println(release)
			fmt.Println(err)
			return internalErrors.ErrorDuplicateName
		}
	}

	return nil
}

//LockRelease lock a release, makes the private project public.
//this is one time operation an can not be reverted
func (s ReleaseStore) LockRelease(releaseID, appID, userID int64) error {
	b := s.Session().Builder()
	q := b.
		Select("releases.id", "releases.app_id", "releases.platform", "releases.note", "releases.version", "releases.created_at", "releases.private").
		From("releases").
		Join("apps").
		On("apps.id=releases.app_id").
		Join("apps_users_permissions").
		On("apps.id=apps_users_permissions.app_id").
		Where("releases.id=? AND apps_users_permissions.user_id=? AND apps_users_permissions.app_id=? AND apps_users_permissions.permission!=?", releaseID, userID, appID, MEMBER)

	var release *Release

	err := q.Iterator().One(&release)

	if err != nil {
		logme.Warn(err.Error())
		return internalErrors.ErrorReleaseNotFound
	}

	if release == nil {
		return internalErrors.ErrorReleaseNotFound
	}

	//once private becomes public, there is no turning back
	if release.Private == false {
		return internalErrors.ErrorReleaseAlreadyLocked
	}

	release.Private = false

	if err = s.Save(release); err != nil {
		return internalErrors.ErrorSomethingWentWrong
	}

	return nil
}
