package users

import (
	"log"

	"github.com/go-pg/pg/v10"
)

type User struct {
	UserId   int64  `sql:"user_id,pk"`
	Username string `sql:"username"`
	Email    string `sql:"email"`
	Password string `sql:"password"`
}

func NewUser() *User {
	return &User{Username: "username1", Email: "email1", Password: "password"}
}

func (u *User) InsertUser(db *pg.DB) {
	_, insertErr := db.Model(u).Insert()
	if insertErr != nil {
		log.Printf("Error inserting user: %v", insertErr)
	}
}
