package main

import (
	"log"
	"net"

	"github.com/gauravgahlot/dockerdoodle/pkg/constants"
	"github.com/gauravgahlot/dockerdoodle/pkg/pb"
	"github.com/gauravgahlot/dockerdoodle/server/services"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen(constants.ServerNetwork, constants.ServerPort)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	// register services with gRPC server
	registerServices(server)

	log.Println("Server listening at port", constants.ServerPort)
	server.Serve(lis)
}

func registerServices(s *grpc.Server) {
	pb.RegisterDockerHostServiceServer(s, new(services.DockerHostService))
	pb.RegisterContainerServiceServer(s, new(services.ContainerService))
}
