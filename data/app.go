package data

import (
	"time"

	"upper.io/bond"
)

//App struct for storing the basic information about each app
type App struct {
	ID         int64     `db:"id,omitempty,pk" json:"-"`
	SecureID   SecureID  `json:"id"`
	Name       string    `db:"name" json:"name"`
	PublicKey  string    `db:"public_key" json:"public_key"`
	PrivateKey string    `db:"private_key" json:"private_key"`
	CreateAt   time.Time `db:"created_at" json:"created_at"`
}

//CollectionName returns collection name in database
func (b *App) CollectionName() string {
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

func (a *App) BeforeCreate(sess bond.Session) error {
	a.ID = int64(a.SecureID)
	return nil
}

func (a *App) AfterCreate(sess bond.Session) {
	a.SecureID = SecureID(a.ID)
}

func (a *App) BeforeUpdate(sess bond.Session) error {
	a.ID = int64(a.SecureID)
	return nil
}

func (a *App) AfterUpdate(sess bond.Session) {
	a.SecureID = SecureID(a.ID)
}

func (a *App) BeforeDelete(sess bond.Session) error {
	a.ID = int64(a.SecureID)
	return nil
}

func (a *App) AfterDelete(sess bond.Session) {
	a.SecureID = SecureID(a.ID)
}
