package data

import (
	"time"

	"upper.io/bond"
)

//AppWithPermission is subset of App which includes Permission
type AppWithPermission struct {
	App                  `bond:",inline"`
	AppsUsersPermissions `bond:",inline"`
}

//App struct for storing the basic information about each app
type App struct {
	ID         int64      `db:"id,omitempty,pk" json:"-"`
	SecureID   SecureID   `json:"id"`
	Name       string     `db:"name" json:"name"`
	PublicKey  string     `db:"public_key" json:"public_key"`
	PrivateKey string     `db:"private_key" json:"private_key,omitempty"`
	CreateAt   *time.Time `db:"created_at" json:"created_at" bondb:",utc"`
}

//CollectionName returns collection name in database
func (a *App) CollectionName() string {
	return `apps`
}

var _ interface {
	bond.HasBeforeCreate
	bond.HasAfterCreate
	bond.HasBeforeUpdate
	bond.HasAfterUpdate
	bond.HasBeforeDelete
	bond.HasAfterDelete
} = &App{}

//BeforeCreate convert Secure ID to regualr id
func (a *App) BeforeCreate(sess bond.Session) error {
	a.ID = int64(a.SecureID)
	return nil
}

//AfterCreate convert regualr id to Secure ID and
//remove private key once it's being created
func (a *App) AfterCreate(sess bond.Session) {
	a.SecureID = SecureID(a.ID)
	a.PrivateKey = ""
}

//BeforeUpdate convert Secure ID to regualr id
func (a *App) BeforeUpdate(sess bond.Session) error {
	a.ID = int64(a.SecureID)
	return nil
}

//AfterUpdate convert regualr id to Secure ID
//remove private key once it's being updated
func (a *App) AfterUpdate(sess bond.Session) {
	a.SecureID = SecureID(a.ID)
	a.PrivateKey = ""
}

//BeforeDelete convert Secure ID to regualr id
func (a *App) BeforeDelete(sess bond.Session) error {
	a.ID = int64(a.SecureID)
	return nil
}

//AfterDelete convert regualr id to Secure ID
func (a *App) AfterDelete(sess bond.Session) {
	a.SecureID = SecureID(a.ID)
}
