package data

import (
	"upper.io/bond"
	"upper.io/db"
)

//UserStore store for release
type UserStore struct {
	bond.Store
}

func (s UserStore) FindByEmailPassword(email, password string) (*User, error) {
	var r *User
	if err := s.Find(db.Cond{"email": email, "password": password}).One(&r); err != nil {
		return nil, err
	}
	return r, nil
}
