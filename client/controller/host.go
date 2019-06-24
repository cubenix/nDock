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
	"github.com/gauravgahlot/dockerdoodle/client/ws"
	"github.com/gauravgahlot/dockerdoodle/pkg/types"
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
	res := viewmodels.HostContainers{SelectedHost: r.URL.Path[6:]}
	res.Hosts = []viewmodels.Host{}
	res.Title = "Host Details"

	var hostIP string
	notFound := true
	for _, s := range *h.hosts {
		res.Hosts = append(res.Hosts, viewmodels.Host{Name: s.Name, IP: s.IP})
		if notFound && strings.EqualFold(r.URL.Path[6:], s.Name) {
			hostIP = s.IP
			notFound = false
		}
	}

	all, running, err := helpers.GetContainers(context.Background(), h.client, hostIP, true)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	res.AllContainers = *all
	res.RunningContainers = *running
	tErr := h.hostTemplate.Execute(w, res)
	if tErr != nil {
		log.Fatal(tErr)
	}
}

func (h host) handleHostContainers(w http.ResponseWriter, r *http.Request) {
	res := viewmodels.HostContainers{SelectedHost: r.URL.Path[17:], RunningContainers: []viewmodels.Container{}}
	res.Hosts = []viewmodels.Host{}
	res.Title = "Containers"

	var hostIP string
	notFound := true
	for _, s := range *h.hosts {
		res.Hosts = append(res.Hosts, viewmodels.Host{Name: s.Name, IP: s.IP})
		if notFound && strings.EqualFold(r.URL.Path[17:], s.Name) {
			hostIP = s.IP
			notFound = false
		}
	}

	all, _, err := helpers.GetContainers(context.Background(), h.client, hostIP, false)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	res.AllContainers = *all
	tErr := h.hostContainerTemplate.Execute(w, res)
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
