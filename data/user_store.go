package data

import (
	"fmt"
	"time"

	"github.com/alinz/releasifier/errors"

	"upper.io/bond"
	"upper.io/db"
)

//UserStore store for release
type UserStore struct {
	bond.Store
}

//FindByEmailPassword find user object based on email and password
func (s UserStore) FindByEmailPassword(email, password string) (*User, error) {
	var r *User
	if err := s.Find(db.Cond{"email": email, "password": password}).One(&r); err != nil {
		return nil, err
	}
	return r, nil
}

//Create create a brand new user
func (s UserStore) Create(fullname, email, password string) (*User, error) {
	tx, err := DB.NewTransaction()
	if err != nil {
		return nil, err
	}
	defer tx.Close()

	user := &User{
		Fullname:  fullname,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().UTC().Truncate(time.Second),
	}

	tx.Save(user)

	fmt.Println(user)

	if err = tx.Commit(); err != nil {
		return nil, errors.ErrorDuplicateEmail
	}

	return user, nil
}
