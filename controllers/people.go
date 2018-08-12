package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// The person Type (more like an object)
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// The Address Type
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

//Globle Variable Type is a Person
var people []Person

// Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//https://golang.org/pkg/net/http/
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(people)
}

// Display a single data
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// var person Person
	for _, item := range people {
		if item.ID == params["id"] {
			// person = item
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	//https://golang.org/pkg/net/http/
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Person{})
}

// create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var onePerson Person
	_ = json.NewDecoder(r.Body).Decode(&onePerson)

	//Get Header
	// headerAuthorization := r.Header.Get("Authorization")

	onePerson.ID = params["id"]
	people = append(people, onePerson)

	w.Header().Set("Content-Type", "application/json")

	//https://golang.org/pkg/net/http/
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(people)
}

// Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1])
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")

	//https://golang.org/pkg/net/http/
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(people)
}

func init() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
}
