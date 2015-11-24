package data

import (
	"fmt"
	"time"
)

//AppWithPermission is subset of App which includes Permission
type AppWithPermission struct {
	App                  `bond:",inline"`
	AppsUsersPermissions `bond:",inline"`
}

//App struct for storing the basic information about each app
type App struct {
	ID         SecureID   `db:"id,omitempty,pk" json:"id"`
	Name       string     `db:"name" json:"name"`
	PublicKey  string     `db:"public_key" json:"public_key"`
	PrivateKey string     `db:"private_key" json:"private_key,omitempty"`
	CreateAt   *time.Time `db:"created_at" json:"created_at" bondb:",utc"`
}

//CollectionName returns collection name in database
func (a *App) CollectionName() string {
	return `apps`
}

func (a App) MarshalDB() (interface{}, error) {
	fmt.Printf("########################MarshalDB %v", a)
	return a, nil
}

func (a *App) UnmarshalDB(interface{}) error {
	fmt.Printf("#######################UnmarshalDB %v", a)
	return nil
}
