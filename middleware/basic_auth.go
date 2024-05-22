package middleware

import (
	"encoding/json"
	"myproject/respond"
	"net/http"
)

var (
	USERNAME = "admin"
	PASSWORD = "admin"
)

func MiddlewareBasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		username, password, ok := r.BasicAuth()

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			response := respond.ErrorAuth{
				Message:    "Gagal autentikasi",
				StatusCode: 401,
				Error: struct {
					Detail string `json:"detail"`
				}{
					Detail: "Kredensial tidak valid atau tidak ada token akses",
				},
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized) // Status Code 401
			json.NewEncoder(w).Encode(response)
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
