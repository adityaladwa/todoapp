package server

import "net/http"

func CreateServer() {
	http.HandleFunc("/", getTodos)
	http.ListenAndServe(":8080", nil)
}

func getTodos(w http.ResponseWriter, r *http.Request) {

}
