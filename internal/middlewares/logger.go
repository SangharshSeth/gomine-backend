package middlewares

import (
	"log"
	"net/http"
)

func ApiLogger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method %s, Path %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
