package data

import (
	"errors"
	"fmt"

	"upper.io/bond"
)

//AppStore store for app
type AppStore struct {
	bond.Store
}

//CreateNewApp once the app is created, whoever create the app is title as owner
func (s AppStore) CreateNewApp(userID int64, appName, publicKey, privateKey string) (*App, error) {
	tx, err := DB.NewTransaction()
	if err != nil {
		return nil, err
	}
	defer tx.Close()

	//we need to pass both public and private keys
	if (publicKey != "" && privateKey == "") ||
		(publicKey == "" && privateKey != "") {
		return nil, errors.New("public and private must be provided together")
	}

	//TODO (Ali): we need to check whether publicKey and privateKey are matched.

	app := &App{
		Name:       appName,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}

	tx.Save(app)

	appUserPermission := &AppsUsersPermissions{
		ID:         0,
		UserID:     userID,
		AppID:      app.ID,
		Permission: OWNER,
	}

	if err := tx.Save(appUserPermission); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("Failed to create new app: %q", err)
	}

	return app, nil
}

//FindAllApps returns all the apps that user has access
func (s AppStore) FindAllApps(userID int64) ([]*AppWithPermission, error) {
	b := s.Session().Builder()
	q := b.
		Select("apps.id", "apps.name", "apps.public_key", "apps.private_key", "apps.created_at", "apps_users_permissions.permission as permission").
		From("apps").
		Join("apps_users_permissions").
		On("apps.id=apps_users_permissions.app_id").
		Where("user_id=?", userID)

	var apps []*AppWithPermission

	err := q.Iterator().All(&apps)

	if err != nil {
		return nil, err
	}

	return apps, nil
}

//FindApp returns a single app that user has access
func (s AppStore) FindApp(appID, userID int64) (*AppWithPermission, error) {
	var app *AppWithPermission

	b := s.Session().Builder()
	q := b.
		Select("apps.id", "apps.name", "apps.public_key", "apps.private_key", "apps.created_at", "apps_users_permissions.permission as permission").
		From("apps").
		Join("apps_users_permissions").
		On("apps.id=apps_users_permissions.app_id").
		Where("user_id=? AND apps.id=?", userID, appID)

	err := q.Iterator().One(&app)

	if err != nil {
		return nil, err
	}

	if app == nil {
		return nil, errors.New("app not found")
	}

	return app, nil
}

func (s AppStore) UpdateApp(appID int64, appName, publicKey, privateKey string, userID int64) error {
	var app *App

	b := s.Session().Builder()
	q := b.
		Select("apps.id as id", "apps.name as name", "apps.public_key as public_key", "apps.private_key as private_key", "apps.created_at as created_at").
		From("apps").
		Join("apps_users_permissions").
		On("apps.id=apps_users_permissions.app_id").
		Where("user_id=? AND apps.id=? AND apps_users_permissions.permission!=?", userID, appID, MEMBER)

	err := q.Iterator().One(&app)

	if err != nil {
		return err
	}

	if app == nil {
		return errors.New("app not found")
	}

	app.Name = appName
	app.PrivateKey = privateKey
	app.PublicKey = publicKey

	s.Save(app)

	return nil
}

func (s AppStore) RemoveApp(appID, userID int64) error {
	tx, err := s.Session().NewTransaction()
	if err != nil {
		return err
	}

	defer tx.Close()

	b := tx.Builder()
	q := b.
		Select("apps.id", "apps.name", "apps.public_key", "apps.private_key", "apps.created_at", "apps_users_permissions.permission as permission").
		From("apps").
		Join("apps_users_permissions").
		On("apps.id=apps_users_permissions.app_id").
		Where("user_id=? AND apps.id=?", userID, appID)

	var app *App
	err = q.Iterator().One(&app)

	if err != nil {
		return err
	}

	if app == nil {
		return errors.New("app not found")
	}

	tx.Delete(app)

	return nil
}

//
//
//
//
//
