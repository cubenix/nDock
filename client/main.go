package main

import (
	"log"
	"net/http"

	"github.com/gauravgahlot/watchdock/client/rpc"
	"github.com/gauravgahlot/watchdock/client/services"

	"google.golang.org/grpc"
)

const (
	serverPort = ":5000"
	clientPort = ":8080"
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
	// create a connection to be used by service clients
	conn, err := grpc.Dial("localhost"+serverPort, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// setup the request handler
	handler.Clients = rpc.InitializeClients(conn)
	registerHandlers()

	// setup the server and start listening
	server := http.Server{
		Addr: "localhost" + clientPort,
	}
	log.Println("Client App listening at port", clientPort)
	log.Fatal(server.ListenAndServe())
}

func registerHandlers() {
	log.Println("registering handlers")
	http.HandleFunc("/containers", handler.Routes["/containers"])
	http.HandleFunc("/container", handler.Routes["/container"])
}
