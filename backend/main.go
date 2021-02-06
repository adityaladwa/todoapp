package main

import (
	db "github.com/adityaladwa/todoapp/db/connect"
	"github.com/adityaladwa/todoapp/server"
)

func main() {
	db.Connect()
	server.SetupServer()
}
