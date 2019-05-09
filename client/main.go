package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gauravgahlot/watchdock/client/services"
)

const (
	constAPIEndpoint    = "0.0.0.0:5000"
	constClientEndpoint = "0.0.0.0:8080"
)

var (
	apiEndpoint    = flag.String("api", constAPIEndpoint, "Endpoint for WatchDock API")
	clientEndpoint = flag.String("client", constClientEndpoint, "Endpoint for WatchDock Client")
)

var handler *services.Handler

func init() {
	flag.Parse()
	handler = services.NewHandler()
}

func main() {
	http.HandleFunc("/hosts", handler.Routes["hosts"])

	// create the server and start listening
	server := http.Server{
		Addr: *clientEndpoint,
	}
	log.Println("Magic happens at:", *clientEndpoint)
	log.Fatal(server.ListenAndServe())
}
