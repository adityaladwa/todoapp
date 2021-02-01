package users

import (
	"log"

	"github.com/go-pg/pg/v10"
)

type User struct {
	UserId   int64  `sql:"user_id,pk" json:"user_id"`
	Username string `sql:"username" json:"user_name"`
	Email    string `sql:"email" json:"email"`
	Password string `sql:"password" json:"password"`
}

func NewUser() *User {
	return &User{Username: "username3", Email: "email3", Password: "password"}
}

func (u *User) InsertUser(db *pg.DB) error {
	_, insertErr := db.Model(u).Insert()
	if insertErr != nil {
		log.Printf("Error inserting user: %v", insertErr)
		return insertErr
	}
	return nil
}

func GetUsers(db *pg.DB) []User {
	var users []User
	selectErr := db.Model(&users).Limit(10).Select()
	if selectErr != nil {
		log.Printf("Error selecting users: %v", selectErr)
	}
	return users
}
