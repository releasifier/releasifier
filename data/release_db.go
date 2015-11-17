package data

import (
	"time"

	"upper.io/bond"
)

//Release structure for data represents in database
type Release struct {
	ID       int64      `db:"id,omitempty,pk" json:"-"`
	SecureID SecureID   `json:"id"`
	Note     string     `db:"note" json:"note"`
	Version  Version    `db:"version" json:"version"`
	CreateAt *time.Time `db:"created_at" json:"created_at" bondb:",utc"`
}

//CollectionName returns collection name in database
func (r *Release) CollectionName() string {
	return `releases`
}

var _ interface {
	bond.HasBeforeCreate
	bond.HasAfterCreate
	bond.HasBeforeUpdate
	bond.HasAfterUpdate
	bond.HasBeforeDelete
	bond.HasAfterDelete
} = &Release{}

//BeforeCreate convert Secure ID to regualr id
func (r *Release) BeforeCreate(sess bond.Session) error {
	r.ID = int64(r.SecureID)
	return nil
}

//AfterCreate convert regualr id to Secure ID
func (r *Release) AfterCreate(sess bond.Session) {
	r.SecureID = SecureID(r.ID)
}

//BeforeUpdate convert Secure ID to regualr id
func (r *Release) BeforeUpdate(sess bond.Session) error {
	r.ID = int64(r.SecureID)
	return nil
}

//AfterUpdate convert regualr id to Secure ID
func (r *Release) AfterUpdate(sess bond.Session) {
	r.SecureID = SecureID(r.ID)
}

//BeforeDelete convert Secure ID to regualr id
func (r *Release) BeforeDelete(sess bond.Session) error {
	r.ID = int64(r.SecureID)
	return nil
}

//AfterDelete convert regualr id to Secure ID
func (r *Release) AfterDelete(sess bond.Session) {
	r.SecureID = SecureID(r.ID)
}
