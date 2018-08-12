package main

import (
	"log"
	"net/http"

	L "./routes"
)

// main function to boot up everything
func main() {
	router := L.NewRouter()
	log.Fatal(http.ListenAndServe(":3881", router))
}
