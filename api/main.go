package main

import (
	"log"
	"net/http"

	"github.com/gauravgahlot/watchdock/api/handlers"
)

func init() {
	// TODO: read the config.json file for Docker hosts
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
