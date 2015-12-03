package data

import "time"

//AppWithPermission is subset of App which includes Permission
type AppWithPermission struct {
	App        `bond:",inline"`
	Permission Permission `db:"permission" json:"permission"`
}

//App struct for storing the basic information about each app
type App struct {
	ID         int64     `db:"id,omitempty,pk" json:"id"`
	Name       string    `db:"name" json:"name"`
	PublicKey  string    `db:"public_key" json:"-"`
	PrivateKey string    `db:"private_key" json:"-"`
	CreatedAt  time.Time `db:"created_at" json:"created_at" bondb:",utc"`
	Private    bool      `db:"private" json:"private"`
}

//CollectionName returns collection name in database
func (a *App) CollectionName() string {
	return `apps`
}
