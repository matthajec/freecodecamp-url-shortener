package middleware

import (
	"fmt"
	"net/http"
)

func CORS(next http.Handler) http.Handler {
	fmt.Println("Added CORS")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "https://www.freecodecamp.org")
		next.ServeHTTP(rw, r)
	})
}
