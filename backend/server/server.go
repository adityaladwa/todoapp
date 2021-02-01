package server

import (
	"log"
	"net/http"

	"github.com/adityaladwa/todoapp/db/users"
	"github.com/go-pg/pg/v10"
)

func GetUsers(db *pg.DB, w http.ResponseWriter, r *http.Request) {
	log.Print("Got a request")
	users.GetUsers(db, w, r)
}
