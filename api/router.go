package api

import (
	"github.com/gorilla/mux"
)

//NewRouter creates and returns a new router will all registered routes.
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", alive).Methods("GET")
	router.HandleFunc("/form", formHandler).Methods("Post")
	return router
}
