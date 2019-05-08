package handlers

import (
	"fmt"
	"net/http"

	"github.com/gauravgahlot/watchdock/api/types"
)

// Handler represents struct with some data and request handlers for incoming HTTP requests
type Handler struct {
	Hosts *[]types.Host
}

// DockerHosts function handles the request for route /hosts
func (h *Handler) DockerHosts(w http.ResponseWriter, r *http.Request) {
	for k, h := range *h.Hosts {
		fmt.Printf("Hostname-%v: %v\n", k, h.Name)
	}

	w.Write([]byte("Docker hosts"))
}
