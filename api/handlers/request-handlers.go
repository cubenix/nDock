package handlers

import "net/http"

// Dashboard function handles the request for route /
func Dashboard(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dashboard goes here"))
}
