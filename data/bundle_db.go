package data

import (
	"time"

	"upper.io/bond"
)

//Bundle structure for each bundle represents in database
type Bundle struct {
	ID        int64     `db:"id,omitempty,pk" json:"-"`
	SecureID  SecureID  `json:"id"`
	ReleaseID int64     `db:"release_id" json:"release_id"`
	Hash      string    `db:"hash" json:"hash"`
	Name      string    `db:"name" json:"name"`
	Type      Type      `db:"type" json:"type"`
	CreatedAt time.Time `db:"cretad_at" json:"created_at"`
}

//CollectionName returns collection name in database
func (b *Bundle) CollectionName() string {
	return `bundles`
}

var _ interface {
	bond.HasBeforeCreate
	bond.HasAfterCreate
	bond.HasBeforeUpdate
	bond.HasAfterUpdate
	bond.HasBeforeDelete
	bond.HasAfterDelete
} = &Bundle{}

//BeforeCreate convert Secure ID to regualr id
func (b *Bundle) BeforeCreate(sess bond.Session) error {
	b.ID = int64(b.SecureID)
	return nil
}

//AfterCreate convert regualr id to Secure ID
func (b *Bundle) AfterCreate(sess bond.Session) {
	b.SecureID = SecureID(b.ID)
}

//BeforeUpdate convert Secure ID to regualr id
func (b *Bundle) BeforeUpdate(sess bond.Session) error {
	b.ID = int64(b.SecureID)
	return nil
}

//AfterUpdate convert regualr id to Secure ID
func (b *Bundle) AfterUpdate(sess bond.Session) {
	b.SecureID = SecureID(b.ID)
}

//BeforeDelete convert Secure ID to regualr id
func (b *Bundle) BeforeDelete(sess bond.Session) error {
	b.ID = int64(b.SecureID)
	return nil
}

//AfterDelete convert regualr id to Secure ID
func (b *Bundle) AfterDelete(sess bond.Session) {
	b.SecureID = SecureID(b.ID)
}
