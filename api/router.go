package api

import (
	"github.com/gorilla/mux"
)

//NewRouter creates and returns a new router will all registered routes.
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", alive).Methods("GET")
	router.HandleFunc("/api/form", formHandler).Methods("POST")
	router.HandleFunc("/api/track", pageViewHandler).Methods("POST")
	router.HandleFunc("/api/sensor", sensorHandler).Methods("POST")
	return router
}
