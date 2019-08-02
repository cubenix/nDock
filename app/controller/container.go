package controller

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"

	vm "github.com/gauravgahlot/dockerdoodle/app/viewmodels"
	"github.com/gauravgahlot/dockerdoodle/pkg/svc"
	"github.com/gauravgahlot/dockerdoodle/pkg/types"
)

type container struct {
	hosts             *[]types.Host
	containerTemplate *template.Template
}

func (c container) registerRoutes() {
	http.HandleFunc("/container/start", c.startContainer)
	http.HandleFunc("/container/stop", c.stopContainer)
	http.HandleFunc("/container/remove", c.removeContainer)
}

func (c container) startContainer(w http.ResponseWriter, r *http.Request) {
	var req vm.ContainerOperationRequest
	decErr := json.NewDecoder(r.Body).Decode(&req)
	if decErr != nil {
		log.Fatal("Invalid Request: ", decErr)
		w.WriteHeader(http.StatusBadRequest)
	}
	var hostIP string
	for _, s := range *c.hosts {
		if strings.EqualFold(req.Host, s.Name) {
			hostIP = s.IP
			break
		}
	}
	err := svc.StartContainer(context.Background(), hostIP, req.ID)
	if err != nil {
		log.Fatal(err)
	}
}

func (c container) stopContainer(w http.ResponseWriter, r *http.Request) {
	var req vm.ContainerOperationRequest
	decErr := json.NewDecoder(r.Body).Decode(&req)
	if decErr != nil {
		log.Fatal("Invalid Request: ", decErr)
		w.WriteHeader(http.StatusBadRequest)
	}
	var hostIP string
	for _, s := range *c.hosts {
		if strings.EqualFold(req.Host, s.Name) {
			hostIP = s.IP
			break
		}
	}
	err := svc.StopContainer(context.Background(), hostIP, req.ID)
	if err != nil {
		log.Fatal(err)
	}
}

func (c container) removeContainer(w http.ResponseWriter, r *http.Request) {
	var req vm.ContainerOperationRequest
	decErr := json.NewDecoder(r.Body).Decode(&req)
	if decErr != nil {
		log.Fatal("Invalid Request: ", decErr)
		w.WriteHeader(http.StatusBadRequest)
	}
	var hostIP string
	for _, s := range *c.hosts {
		if strings.EqualFold(req.Host, s.Name) {
			hostIP = s.IP
			break
		}
	}
	err := svc.RemoveContainer(context.Background(), hostIP, req.ID)
	if err != nil {
		log.Fatal(err)
	}
}
