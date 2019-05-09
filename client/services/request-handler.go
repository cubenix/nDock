package services

import (
	"net/http"
)

type handleFunc func(w http.ResponseWriter, r *http.Request)

// Handler represents struct with some data and request handlers for incoming HTTP requests
type Handler struct {
	Routes map[string]handleFunc
}

// NewHandler initializes the request handler and returns a pointer to it
func NewHandler() *Handler {
	h := Handler{}
	h.initializeRoutes()
	return &h
}

func (h *Handler) initializeRoutes() {
	h.Routes = map[string]handleFunc{
		"hosts": h.dockerHosts,
	}
}

func (h *Handler) dockerHosts(w http.ResponseWriter, r *http.Request) {
	// TODO: get data from API
	w.Write([]byte("Welcome to client"))
}
