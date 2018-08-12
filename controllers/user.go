package controllers

import (
	"net/http"
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
)

// Display a single data
func GetUser(w http.ResponseWriter, r *http.Request) {
	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	// defer db.Close()

	// w.Header().Set("Content-Type", "application/json")
	// //https://golang.org/pkg/net/http/
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(&Person{})
}
