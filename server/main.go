package main

import (
	"log"
	"net"

	"github.com/gauravgahlot/dockerdoodle/pb"
	"github.com/gauravgahlot/dockerdoodle/server/services"
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

	// register services with gRPC server
	registerServices(server)

	log.Println("Server listening at port", port)
	server.Serve(lis)
}

func registerServices(s *grpc.Server) {
	pb.RegisterDockerHostServiceServer(s, new(services.DockerHostService))
	pb.RegisterContainerServiceServer(s, new(services.ContainerService))
}
