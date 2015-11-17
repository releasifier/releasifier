package data

import (
	"fmt"

	"upper.io/bond"
)

//AppStore store for app
type AppStore struct {
	bond.Store
}

func (s AppStore) CreateNewApp(appName string) (*App, error) {
	tx, err := DB.NewTransaction()
	if err != nil {
		return nil, err
	}
	defer tx.Close()

	app := &App{
		ID:         0,
		Name:       appName,
		PublicKey:  "",
		PrivateKey: "",
	}

	if err := tx.Save(app); err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("Failed to create new app: %q", err)
	}

	app.SecureID = SecureID(app.ID)

	return app, nil
}
