package services

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gauravgahlot/watchdock/client/rpc"
	"github.com/gauravgahlot/watchdock/client/services/helpers"
	"github.com/gauravgahlot/watchdock/types"
	vm "github.com/gauravgahlot/watchdock/types/viewmodels"
)

type handleFunc func(w http.ResponseWriter, r *http.Request)

type containersReqData struct {
	All bool `json:"all"`
}

// Handler represents struct with some data and request handlers for incoming HTTP requests
type Handler struct {
	hosts     *[]types.Host
	Clients   *rpc.Clients
	Templates map[string]*template.Template
	Routes    map[string]handleFunc
}

// NewHandler initializes the request handler and returns a pointer to it
func NewHandler(hosts *[]types.Host) *Handler {
	h := Handler{hosts: hosts}
	h.initializeRoutes()
	return &h
}

func (h *Handler) initializeRoutes() {
	h.Routes = map[string]handleFunc{
		"home":       h.home,
		"containers": h.containers,
	}
}

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	requestedFile := r.URL.Path[1:]
	template := h.Templates[requestedFile+".html"]
	if template != nil {
		res, re := helpers.GetContainersCount(h.Clients.DockerService, h.hosts, false)
		if re != nil {
			log.Fatal(re)
			w.WriteHeader(http.StatusInternalServerError)
		}

		context := vm.Home{Hosts: *res}
		context.Title = "Home"
		err := template.Execute(w, context)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *Handler) containers(w http.ResponseWriter, r *http.Request) {
	var data containersReqData
	decErr := json.NewDecoder(r.Body).Decode(&data)
	if decErr != nil {
		log.Fatal("Error receiving data: ", decErr)
		w.WriteHeader(http.StatusBadRequest)
	}

	res, err := helpers.GetContainersCount(h.Clients.DockerService, h.hosts, data.All)
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
