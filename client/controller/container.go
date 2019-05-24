package controller

import (
	"html/template"
	"net/http"

	"github.com/gauravgahlot/dockerdoodle/client/rpc"
	"github.com/gauravgahlot/dockerdoodle/types"
)

type container struct {
	hosts             *[]types.Host
	containerTemplate *template.Template
	client            rpc.ContainerServiceClient
}

func (c container) registerRoutes() {
	http.HandleFunc("/container/", c.handleContainerDetails)
}

func (c container) handleContainerDetails(w http.ResponseWriter, r *http.Request) {
	host := r.FormValue("host")
	http.Redirect(w, r, "/host/containers/"+host, http.StatusPermanentRedirect)
}
