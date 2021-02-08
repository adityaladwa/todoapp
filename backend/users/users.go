package users

import (
	"log"

	db "github.com/adityaladwa/todoapp/db/connect"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId   int64  `sql:"user_id,pk" json:"user_id"`
	Username string `sql:"username" json:"user_name"`
	Email    string `sql:"email" json:"email"`
	Password string `sql:"password" json:"password"`
}

func NewUser(username string, email string, password string) *User {
	hash := GetPasswordHash(password)
	return &User{Username: username, Email: email, Password: string(hash)}
}

func (u *User) InsertUser() error {
	u.Password = string(GetPasswordHash(u.Password))
	_, insertErr := db.DBConn.Model(u).Insert()
	if insertErr != nil {
		log.Printf("Error inserting user: %v", insertErr)
		return insertErr
	}
	return nil
}

func GetUsers() []User {
	var users []User
	selectErr := db.DBConn.Model(&users).Limit(10).Select()
	if selectErr != nil {
		log.Printf("Error selecting users: %v", selectErr)
	}
	return users
}

func GetPasswordHash(password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error generating password hash: %v", err)
	}
	return hash
}

func IsPasswordValid(passwordHash string, password string) bool {
	decryptErr := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if decryptErr != nil {
		log.Printf("Error decrypting password: %v", decryptErr)
	}
	log.Printf("Password was decrypted successfully")
	return true
}
