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
		ID:         0,
		Name:       appName,
		PublicKey:  "",
		PrivateKey: "",
	}

	if err := tx.Save(app); err != nil {
		return nil, err
	}

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

	app.SecureID = SecureID(app.ID)

	return app, nil
}

func (s AppStore) FindAllByUserID(userID int64) ([]*AppWithPermission, error) {
	var apps []*AppWithPermission

	q := squirrel.
		Select("apps.id as id, apps.name as name, apps.public_key as public_key, apps.private_key as private_key, apps.created_at as created_at, apps_users_permissions.permission as permission").
		From("apps LEFT JOIN apps_users_permissions ON apps.id=apps_users_permissions.app_id").
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

	return apps, nil
}
