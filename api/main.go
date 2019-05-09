package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gauravgahlot/watchdock/api/services"
)

const constAPIEndpoint = "0.0.0.0:5000"

var apiEndpoint = flag.String("api", constAPIEndpoint, "Endpoint for WatchDock API")
var handler *services.Handler

func init() {
	flag.Parse()
	var reader services.ConfigReader = services.JSONConfigReader{}
	conf, err := reader.ReadConfig()

	if err != nil {
		log.Fatalln("Failed to read the configuration")
	}
	handler = services.NewHandler(&conf.Hosts)
}

func main() {
	http.HandleFunc("/hosts", handler.Routes["hosts"])

	// create the server and start listening
	server := http.Server{
		Addr: *apiEndpoint,
	}
	log.Println("Server Listening at:", *apiEndpoint)
	log.Fatal(server.ListenAndServe())
}
