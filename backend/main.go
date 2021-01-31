package main

import (
	"log"
	"net/http"

	db "github.com/adityaladwa/todoapp/db/connect"
	"github.com/adityaladwa/todoapp/db/users"
	"github.com/go-pg/pg/v10"
)

type App struct {
	DB *pg.DB
}

func (a *App) Init() error {
	conn := db.Connect()
	a.DB = conn
	return nil
}

func main() {
	app := &App{}
	app.Init()
	defer app.DB.Close()
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Got a request")
		users.GetUsers(app.DB, w, r)
	})
	http.ListenAndServe(":8080", nil)
}
