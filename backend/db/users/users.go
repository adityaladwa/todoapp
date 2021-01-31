package users

import (
	"encoding/json"
	"log"
	"net/http"

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

func GetUsers(conn *pg.DB, w http.ResponseWriter, r *http.Request) {
	var users []User
	selectErr := conn.Model(&users).Limit(10).Select()
	if selectErr != nil {
		log.Printf("Error selecting users: %v", selectErr)
	}
	usersJson, err := json.Marshal(users)
	if err != nil {
		log.Printf(err.Error())
	}
	w.Write(usersJson)
}
