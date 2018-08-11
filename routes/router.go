package routes

import "github.com/gorilla/mux"

//New Router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	peopleRouter(r)
	homeRouter(r)
	return r
}
