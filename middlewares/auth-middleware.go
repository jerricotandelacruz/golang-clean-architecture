package middleware

import (
	"fmt"
	"net/http"
)

func AuthenticationMiddleware() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			fmt.Println("Authenticating...")

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}
