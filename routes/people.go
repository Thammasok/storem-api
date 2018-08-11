package routes

import (
	PeopleController "../controllers"
	"github.com/gorilla/mux"
)

func peopleRouter(router *mux.Router) {
	router.HandleFunc("/people", PeopleController.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", PeopleController.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", PeopleController.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", PeopleController.DeletePerson).Methods("DELETE")
}
