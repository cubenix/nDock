package main

import (
	"context"
	"log"

	"github.com/gauravgahlot/watchdock/client/rpc"
	"github.com/gauravgahlot/watchdock/client/services"
	"github.com/gauravgahlot/watchdock/pb"

	"google.golang.org/grpc"
)

const (
	serverPort = ":5000"
	clientPort = ":8080"
)

var (
	handler *services.Handler
	clients *rpc.Clients
)

func init() {
	var reader services.ConfigReader = services.JSONConfigReader{}
	conf, err := reader.ReadConfig()

	if err != nil {
		log.Fatalln("Failed to read the configuration")
	}
	handler = services.NewHandler(&conf.Hosts)
}

func main() {
	conn, err := grpc.Dial("localhost"+serverPort, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	clients = rpc.InitializeClients(conn)
	// server := http.Server{
	// 	Addr: "localhost" + clientPort,
	// }
	// log.Println("Client App listening at port", clientPort)
	// log.Fatal(server.ListenAndServe())

	getContainersCount(clients.DockerService)
	getContainer(clients.ContainerService)
}

func getContainersCount(client pb.DockerHostServiceClient) {
	res, err := client.GetContainersCount(context.Background(), &pb.GetContainersCountRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.HostContainers)
}

func getContainer(client pb.ContainerServiceClient) {
	res, err := client.GetContainer(context.Background(), &pb.GetContainerRequest{Id: "container-id"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
