package api

import (
	"net/http"
)

// Auth ...
func Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	header, ok := r.Header["X-Auth-Token"]
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("not authenticated"))
	}

	// Check Token exist
	if !isTokenExist(header[0]) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("invalid credentials"))
	}

	user := getUser(header[0])

	w.Header().Set("X-Auth-Token", header[0])
	w.Header().Set("X-User-Id", user)
	w.Header().Set("X-Service", "app")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("authenticated"))
}

func isTokenExist(token string) bool {
	return true
}

func getUser(token string) string {
	return "123"
}
