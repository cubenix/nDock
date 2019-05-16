package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gauravgahlot/dockerdoodle/client/rpc"
	"github.com/gauravgahlot/dockerdoodle/client/viewmodels"
	"github.com/gauravgahlot/dockerdoodle/types"
)

type host struct {
	hostTemplate *template.Template
	hosts        *[]types.Host
	client       rpc.ContainerServiceClient
}

func (h host) registerRoutes() {
	// http.HandleFunc("/hosts", h.handleHosts)
	http.HandleFunc("/host/", h.handleHosts)
}

func (h host) handleHosts(w http.ResponseWriter, req *http.Request) {
	log.Println("made it to the controller: ", req.URL.Path)

	context := viewmodels.DockerHost{Hosts: []viewmodels.Host{}}
	context.Title = "Hosts"
	for _, s := range *h.hosts {
		context.Hosts = append(context.Hosts, viewmodels.Host{Name: s.Name, IP: s.IP})
	}

	err := h.hostTemplate.Execute(w, context)
	if err != nil {
		log.Fatal(err)
	}
}
