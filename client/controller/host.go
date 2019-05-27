package controller

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gauravgahlot/dockerdoodle/client/helpers"
	"github.com/gauravgahlot/dockerdoodle/client/rpc"
	"github.com/gauravgahlot/dockerdoodle/client/viewmodels"
	"github.com/gauravgahlot/dockerdoodle/client/ws"
	"github.com/gauravgahlot/dockerdoodle/types"
)

type host struct {
	hostTemplate          *template.Template
	hostContainerTemplate *template.Template
	hosts                 *[]types.Host
	client                rpc.DockerServiceClient
}

func (h host) registerRoutes() {
	http.HandleFunc("/host/", h.handleHosts)
	http.HandleFunc("/host/containers/", h.handleHostContainers)
	http.HandleFunc("/ws", wsEndpoint)
}

func (h host) handleHosts(w http.ResponseWriter, r *http.Request) {
	context := viewmodels.HostContainers{SelectedHost: r.URL.Path[6:]}
	context.Hosts = []viewmodels.Host{}
	context.Title = "Host Details"

	var hostIP string
	notFound := true
	for _, s := range *h.hosts {
		context.Hosts = append(context.Hosts, viewmodels.Host{Name: s.Name, IP: s.IP})
		if notFound && strings.EqualFold(r.URL.Path[6:], s.Name) {
			hostIP = s.IP
			notFound = false
		}
	}

	all, running, err := helpers.GetContainers(h.client, hostIP, true)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	context.AllContainers = *all
	context.RunningContainers = *running
	tErr := h.hostTemplate.Execute(w, context)
	if tErr != nil {
		log.Fatal(tErr)
	}
}

func (h host) handleHostContainers(w http.ResponseWriter, r *http.Request) {
	context := viewmodels.HostContainers{SelectedHost: r.URL.Path[17:], RunningContainers: []viewmodels.Container{}}
	context.Hosts = []viewmodels.Host{}
	context.Title = "Containers"

	var hostIP string
	notFound := true
	for _, s := range *h.hosts {
		context.Hosts = append(context.Hosts, viewmodels.Host{Name: s.Name, IP: s.IP})
		if notFound && strings.EqualFold(r.URL.Path[17:], s.Name) {
			hostIP = s.IP
			notFound = false
		}
	}
	all, _, err := helpers.GetContainers(h.client, hostIP, false)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	context.AllContainers = *all
	tErr := h.hostContainerTemplate.Execute(w, context)
	if tErr != nil {
		log.Fatal(tErr)
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	hub := ws.NewHub()
	go hub.Run()
	ws.ServeWs(hub, w, r)
	helpers.Hub = hub
}
