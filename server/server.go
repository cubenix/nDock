package server

import (
	"log"
	"net"

	"github.com/gauravgahlot/dockerdoodle/constants"
	"github.com/gauravgahlot/dockerdoodle/pb"
	"github.com/gauravgahlot/dockerdoodle/server/services"
	"google.golang.org/grpc"
)

// Initialize creates a gRPC server that interacts with the Docker hosts
func Initialize() {
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
