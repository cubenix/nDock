package controller

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gauravgahlot/dockerdoodle/client/helpers"
	"github.com/gauravgahlot/dockerdoodle/client/rpc"
	vm "github.com/gauravgahlot/dockerdoodle/client/viewmodels"
	"github.com/gauravgahlot/dockerdoodle/pkg/types"
)

type container struct {
	hosts             *[]types.Host
	containerTemplate *template.Template
	client            rpc.ContainerServiceClient
}

func (c container) registerRoutes() {
	http.HandleFunc("/container/start", c.startContainer)
	http.HandleFunc("/container/stop", c.stopContainer)
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
	err := helpers.StartContainer(context.Background(), c.client, hostIP, req.ID)
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
	err := helpers.StopContainer(context.Background(), c.client, hostIP, req.ID)
	if err != nil {
		log.Fatal(err)
	}
}
