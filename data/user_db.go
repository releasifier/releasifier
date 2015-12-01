package data

import "time"

//User struct for storing the basic information about each user
type User struct {
	ID        int64     `db:"id,omitempty,pk" json:"id"`
	Fullname  string    `db:"fullname" json:"fullname"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at" bondb:",utc"`
}

//CollectionName returns collection name in database
func (u *User) CollectionName() string {
	return `users`
}
