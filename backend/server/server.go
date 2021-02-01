package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adityaladwa/todoapp/db/users"
	"github.com/go-pg/pg/v10"
)

func UsersEndpoint(db *pg.DB, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(db, w, r)
	case http.MethodPost:
		addUser(db, w, r)
	default:
		log.Printf("Only get and post is available")
	}
}

func getUsers(db *pg.DB, w http.ResponseWriter, r *http.Request) {
	log.Print("Got a request")
	userResult := users.GetUsers(db)
	usersJson, err := json.Marshal(userResult)
	if err != nil {
		log.Printf(err.Error())
	}
	w.Write(usersJson)
}

func addUser(db *pg.DB, w http.ResponseWriter, r *http.Request) {
	log.Print("Got a request to add a user")
	user := users.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	defer r.Body.Close()
	if err := user.InsertUser(db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	response, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(response))
}
