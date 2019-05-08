package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gauravgahlot/watchdock/api/handlers"
	"github.com/gauravgahlot/watchdock/api/services"
)

func init() {
	var reader services.ConfigReader = services.JSONConfigReader{}
	conf, err := reader.ReadConfig()

	if err != nil {
		log.Fatalln("Failed to read the configuration")
	}

	for k, h := range conf.Hosts {
		fmt.Printf("Hostname-%v: %v\n", k, h.Name)
	}
}

func main() {
	http.HandleFunc("/", handlers.Dashboard)

	// create the server and start listening
	server := http.Server{
		Addr: "0.0.0.0:5000",
	}
	log.Println("Server Listening at PORT: 5000")
	log.Fatal(server.ListenAndServe())
}
