package main

import (
	"github.com/adityaladwa/todoapp/cache"
	db "github.com/adityaladwa/todoapp/db/connect"
	"github.com/adityaladwa/todoapp/server"
)

func main() {
	db.Connect()
	defer db.DBConn.Close()
	cache.Connect()
	server.SetupServer()
}
