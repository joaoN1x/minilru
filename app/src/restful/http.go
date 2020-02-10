package restful

import (
	"encoding/json"
	"net/http"
)

// answer rest format
func writeOut(w http.ResponseWriter, r *http.Request, code int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}

// rest middleware handling
func MiddleOne(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

// some auth could be goin on here
func BasicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// could do Basic Auth
		// or check JWT token

		h.ServeHTTP(w, r)
	}
}
