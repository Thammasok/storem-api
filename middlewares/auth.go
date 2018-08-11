package middlewares

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
)

// for use on route (using a http.HandlerFunc)
func AuthBasic(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read basic auth information
		usr, _, ok := r.BasicAuth()

		// if there is no basic auth (no matter which credentials)
		if !ok {
			errMsg := "Authentication error!"
			// return a 403 forbidden
			http.Error(w, errMsg, http.StatusForbidden)
			log.Println(errMsg)

			// stop processing route
			return
		}

		// let's assume we check the user against a database to get
		// his admin-right and put this to the request context
		context.Set(r, "isAdmin", true)

		// else continue processing
		log.Printf("User %s logged in.", usr)
		next(w, r)
	}
}
