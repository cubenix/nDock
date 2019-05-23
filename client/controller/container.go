package controller

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gauravgahlot/dockerdoodle/client/helpers"
	"github.com/gauravgahlot/dockerdoodle/client/rpc"
	"github.com/gauravgahlot/dockerdoodle/client/viewmodels"
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
	var hostIP string
	for _, h := range *c.hosts {
		if strings.EqualFold(h.Name, r.FormValue("host")) {
			hostIP = h.IP
			break
		}
	}
	res, err := helpers.GetContainer(context.Background(), c.client, hostIP, r.FormValue("containerID"))
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusNotFound)
	}

	ctx := viewmodels.ContainerDetails{}
	ctx.Container = *res
	ctx.Title = "Container Details"
	c.containerTemplate.Execute(w, ctx)
}
