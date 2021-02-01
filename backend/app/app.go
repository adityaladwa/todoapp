package app

import (
	"log"
	"net/http"

	db "github.com/adityaladwa/todoapp/db/connect"
	"github.com/adityaladwa/todoapp/server"
	"github.com/go-pg/pg/v10"
)

type App struct {
	DB *pg.DB
}

func (app *App) Init() error {
	conn := db.Connect()
	app.DB = conn
	app.setupRoutes()
	return nil
}

func (app *App) handleRequest(handler func(db *pg.DB, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.DB, w, r)
	}
}

func (app *App) setupRoutes() {
	log.Printf("Setting up routes")
	http.HandleFunc("/users", app.handleRequest(server.UsersEndpoint))
	log.Printf("Server is listening on 8080")
	http.ListenAndServe(":8080", nil)
}
