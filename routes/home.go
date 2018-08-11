package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// The home Type (more like an object)
type Home struct {
	TITLE string `json:"title,omitempty"`
}

//Globle Variable Type is a Person
var home []Home

// Display home
func GetHome(w http.ResponseWriter, r *http.Request) {
	//https://golang.org/pkg/net/http/
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	home = append(home, Home{TITLE: "Home page"})

	json.NewEncoder(w).Encode(home)
}

func homeRouter(router *mux.Router) {
	router.HandleFunc("/", GetHome).Methods("GET")
}
