package routes

import (
	HomeController "../controllers"
	Middlewares "../middlewares"
	"github.com/gorilla/mux"
)

func homeRouter(router *mux.Router) {
	router.HandleFunc("/", Middlewares.AuthBasic(HomeController.GetHome)).Methods("GET")
}
