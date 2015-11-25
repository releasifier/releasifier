package data

import (
	"errors"
	"fmt"

	"github.com/lann/squirrel"

	"upper.io/bond"
	"upper.io/db"
	"upper.io/db/util/sqlutil"
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

	app := &App{
		Name:       appName,
		PublicKey:  "",
		PrivateKey: "",
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
	var apps []*AppWithPermission

	q := squirrel.
		Select(`apps.id as id,
						apps.name as name,
						apps.public_key as public_key,
						apps.private_key as private_key,
						apps.created_at as created_at,
						apps_users_permissions.permission as permission`).
		From(`apps LEFT JOIN apps_users_permissions ON apps.id=apps_users_permissions.app_id`).
		Where(squirrel.Eq(db.Cond{"user_id": userID}))

	sql, args, err := q.PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := DB.Sqlx.Queryx(sql, args...)
	if err != nil {
		return nil, err
	}

	if err = sqlutil.FetchRows(rows, &apps); err != nil {
		return nil, err
	}

	if apps == nil || len(apps) == 0 {
		return make([]*AppWithPermission, 0), nil
	}

	return apps, nil
}

//FindApp returns a single app that user has access
func (s AppStore) FindApp(appID, userID int64) (*AppWithPermission, error) {
	var apps []*AppWithPermission

	cond := squirrel.And{
		squirrel.Eq{"user_id": userID},
		squirrel.Eq{"apps.id": appID},
	}

	q := squirrel.
		Select(`apps.id as id,
						apps.name as name,
						apps.public_key as public_key,
						apps.private_key as private_key,
						apps.created_at as created_at,
						apps_users_permissions.permission as permission`).
		From(`apps LEFT JOIN apps_users_permissions ON apps.id=apps_users_permissions.app_id`).
		Where(cond)

	sql, args, err := q.PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := DB.Sqlx.Queryx(sql, args...)
	if err != nil {
		return nil, err
	}

	if err = sqlutil.FetchRows(rows, &apps); err != nil {
		return nil, err
	}

	if len(apps) == 1 {
		return apps[0], nil
	}

	return nil, errors.New("app not found")
}
