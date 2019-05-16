package controller

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gauravgahlot/dockerdoodle/client/helpers"
	"github.com/gauravgahlot/dockerdoodle/client/rpc"
	"github.com/gauravgahlot/dockerdoodle/client/viewmodels"
	vm "github.com/gauravgahlot/dockerdoodle/client/viewmodels"
	"github.com/gauravgahlot/dockerdoodle/constants"
	"github.com/gauravgahlot/dockerdoodle/types"
)

type home struct {
	homeTemplate *template.Template
	client       rpc.DockerServiceClient
	hosts        *[]types.Host
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/containers-count", h.handleContainersCount)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	context := vm.Home{Hosts: []viewmodels.Host{}}
	context.Title = "Home"

	for i, host := range *h.hosts {
		context.Hosts = append(context.Hosts, viewmodels.Host{
			Name:        host.Name,
			IP:          host.IP,
			BGColor:     constants.BGClasses[i],
			ColorCode:   constants.TextClasses[i],
			BGColorCode: constants.BGCodes[i],
		})
	}
	err := h.homeTemplate.Execute(w, context)
	if err != nil {
		log.Fatal(err)
	}
}

func (h home) handleContainersCount(w http.ResponseWriter, r *http.Request) {
	var data vm.ContainersCountRequest
	decErr := json.NewDecoder(r.Body).Decode(&data)
	if decErr != nil {
		log.Fatal("Error receiving data: ", decErr)
		w.WriteHeader(http.StatusBadRequest)
	}

	res, err := helpers.GetContainersCount(h.client, h.hosts, data.All)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	encErr := json.NewEncoder(w).Encode(vm.Home{Hosts: *res})
	if encErr != nil {
		log.Fatal("Error sending data: ", encErr)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
