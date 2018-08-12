package routes

import (
	Controller "../controllers"
	Middlewares "../middlewares"
	"github.com/gorilla/mux"
)

func home(router *mux.Router) {
	router.HandleFunc("/", Middlewares.AuthBasic(Controller.GetHome)).Methods("GET")
}
