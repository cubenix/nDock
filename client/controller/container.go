package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gauravgahlot/dockerdoodle/client/rpc"
)

type container struct {
	containerTemplate *template.Template
	client            rpc.ContainerServiceClient
}

func (c container) registerRoutes() {
	http.HandleFunc("/container/", c.handleContainerDetails)
}

func (c container) handleContainerDetails(w http.ResponseWriter, r *http.Request) {
	containerID := r.FormValue("containerID")
	log.Println(containerID)

	w.Write([]byte("<h1>Container Details Page</h1>"))
}
