package services

import (
	"encoding/json"
	"net/http"

	"github.com/gauravgahlot/watchdock/types"
)

type handleFunc func(w http.ResponseWriter, r *http.Request)

// Handler represents struct with some data and request handlers for incoming HTTP requests
type Handler struct {
	Hosts  *[]types.Host
	Routes map[string]handleFunc
}

// NewHandler initializes the request handler and returns a pointer to it
func NewHandler(hosts *[]types.Host) *Handler {
	h := Handler{Hosts: hosts}
	h.initializeRoutes()
	return &h
}

func (h *Handler) initializeRoutes() {
	h.Routes = map[string]handleFunc{
		"hosts": h.dockerHosts,
	}
}

func (h *Handler) dockerHosts(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(h.Hosts)
	w.Write([]byte(data))
}
