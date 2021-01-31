package users

import (
	"log"

	"github.com/go-pg/pg/v10"
)

type User struct {
	UserId   int64  `sql:"user_id,pk",json:"user_id"`
	Username string `sql:"username",json:"user_name`
	Email    string `sql:"email",json:"email`
	Password string `sql:"password",json:"password`
}

func NewUser() *User {
	return &User{Username: "username3", Email: "email3", Password: "password"}
}

func (u *User) InsertUser(conn *pg.DB) {
	_, insertErr := conn.Model(u).Insert()
	if insertErr != nil {
		log.Printf("Error inserting user: %v", insertErr)
	}
}

func GetUsers(conn *pg.DB) []User {
	var users []User
	selectErr := conn.Model(&users).Select()
	if selectErr != nil {
		log.Printf("Error selecting users: %v", selectErr)
	}
	return users
}
