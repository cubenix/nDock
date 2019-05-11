package main

import (
	"log"
	"net"

	"github.com/gauravgahlot/watchdock/pb"
	"github.com/gauravgahlot/watchdock/server/services"
	"google.golang.org/grpc"
)

const (
	port    = ":5000"
	network = "tcp"
)

func main() {
	lis, err := net.Listen(network, port)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterDockerHostServiceServer(server, new(services.DockerHostService))
	pb.RegisterContainerServiceServer(server, new(services.ContainerService))
	log.Println("Server listening at port", port)
	server.Serve(lis)
}
