package main

import (
	"encoding/json"
	"log"

	db "github.com/adityaladwa/todoapp/db/connect"
	"github.com/adityaladwa/todoapp/db/users"
)

func main() {
	conn := db.Connect()
	defer conn.Close()
	users := users.GetUsers(conn)
	usersJson, err := json.Marshal(users)
	if err != nil {
		log.Printf(err.Error())
	}
	log.Printf("Users json %v", string(usersJson))
}
