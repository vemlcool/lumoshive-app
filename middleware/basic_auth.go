package middleware

import "net/http"

var (
	USERNAME = "admin"
	PASSWORD = "admin"
)

func MiddlewareBasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		username, password, ok := r.BasicAuth()

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error Unauthorized"))
			return
		}

		isValid := (username == USERNAME && PASSWORD == password)
		if !isValid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error Unauthorized"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
