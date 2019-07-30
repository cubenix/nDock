package controller

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gauravgahlot/dockerdoodle/app/svc"
	"github.com/gauravgahlot/dockerdoodle/client/viewmodels"
	vm "github.com/gauravgahlot/dockerdoodle/client/viewmodels"
	"github.com/gauravgahlot/dockerdoodle/pkg/constants"
	cnv "github.com/gauravgahlot/dockerdoodle/pkg/converters"
	"github.com/gauravgahlot/dockerdoodle/pkg/types"
)

type home struct {
	homeTemplate *template.Template
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

	res, err := svc.GetContainersCount(context.Background(), h.hosts, data.All)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	encErr := json.NewEncoder(w).Encode(vm.Home{Hosts: *cnv.ToHostsViewModel(*res, *h.hosts)})
	if encErr != nil {
		log.Fatal("Error sending data: ", encErr)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
