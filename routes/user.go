package routes

import (
	Controller "../controllers"
	"github.com/gorilla/mux"
)

func user(router *mux.Router) {
	router.HandleFunc("/user", Controller.GetUser).Methods("GET")
}
