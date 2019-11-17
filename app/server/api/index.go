package api

import (
	"fmt"
	"net/http"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	auth := r.Header["X-Auth-Token"][0]
	user := r.Header["X-User-Id"][0]
	service := r.Header["X-Service"][0]

	// bytes, _ := json.Marshal(r.Header)

	returnInfo := fmt.Sprintf(
		"Index Page for user: %s, service: %s with auth: %s",
		user, service, auth)

	w.Write([]byte(returnInfo))
}
