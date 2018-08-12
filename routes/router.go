package routes

import "github.com/gorilla/mux"

//New Router
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	people(r)
	home(r)
	user(r)

	return r
}
