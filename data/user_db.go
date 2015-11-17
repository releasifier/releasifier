package data

import (
	"time"

	"upper.io/bond"
)

//User struct for storing the basic information about each user
type User struct {
	ID       int64      `db:"id,omitempty,pk" json:"-"`
	SecureID SecureID   `json:"id"`
	Fullname string     `db:"fullname" json:"fullname"`
	Email    string     `db:"email" json:"email"`
	Password string     `db:"password" json:"password"`
	CreateAt *time.Time `db:"created_at" json:"created_at" bondb:",utc"`
}

//CollectionName returns collection name in database
func (u *User) CollectionName() string {
	return `users`
}

var _ interface {
	bond.HasBeforeCreate
	bond.HasAfterCreate
	bond.HasBeforeUpdate
	bond.HasAfterUpdate
	bond.HasBeforeDelete
	bond.HasAfterDelete
} = &User{}

//BeforeCreate convert Secure ID to regualr id
func (u *User) BeforeCreate(sess bond.Session) error {
	u.ID = int64(u.SecureID)
	return nil
}

//AfterCreate convert regualr id to Secure ID
func (u *User) AfterCreate(sess bond.Session) {
	u.SecureID = SecureID(u.ID)
}

//BeforeUpdate convert Secure ID to regualr id
func (u *User) BeforeUpdate(sess bond.Session) error {
	u.ID = int64(u.SecureID)
	return nil
}

//AfterUpdate convert regualr id to Secure ID
func (u *User) AfterUpdate(sess bond.Session) {
	u.SecureID = SecureID(u.ID)
}

//BeforeDelete convert Secure ID to regualr id
func (u *User) BeforeDelete(sess bond.Session) error {
	u.ID = int64(u.SecureID)
	return nil
}

//AfterDelete convert regualr id to Secure ID
func (u *User) AfterDelete(sess bond.Session) {
	u.SecureID = SecureID(u.ID)
}
