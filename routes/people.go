package routes

import (
	Controller "../controllers"
	"github.com/gorilla/mux"
)

func people(router *mux.Router) {
	router.HandleFunc("/people", Controller.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", Controller.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", Controller.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", Controller.DeletePerson).Methods("DELETE")
}
