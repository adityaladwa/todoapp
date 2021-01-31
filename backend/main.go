package main

import (
	db "github.com/adityaladwa/todoapp/db/connect"
	"github.com/adityaladwa/todoapp/db/users"
)

func main() {
	db := db.Connect()
	user := users.NewUser()
	user.InsertUser(db)
	defer db.Close()
}
