package main

import (
	"log"
	"net/http"

	"github.com/gauravgahlot/watchdock/api/services"
)

var handler *services.Handler

func init() {
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
		Addr: "0.0.0.0:5000",
	}
	log.Println("Server Listening at PORT: 5000")
	log.Fatal(server.ListenAndServe())
}
