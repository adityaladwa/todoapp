package main

import (
	"log"
	"net/http"

	"github.com/adityaladwa/todoapp/app"
	"github.com/adityaladwa/todoapp/db/users"
)

func main() {
	app := &app.App{}
	app.Init()
	defer app.DB.Close()
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Got a request")
		users.GetUsers(app.DB, w, r)
	})
	http.ListenAndServe(":8080", nil)
}
