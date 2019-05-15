package rpc

import (
	"github.com/gauravgahlot/dockerdoodle/pb"
	"google.golang.org/grpc"
)

// Clients struct to hold all the service clients
type Clients struct {
	DockerService    pb.DockerHostServiceClient
	ContainerService pb.ContainerServiceClient
}

// InitializeClients initializes client for respective service
func InitializeClients(conn *grpc.ClientConn) *Clients {
	c := Clients{
		DockerService:    pb.NewDockerHostServiceClient(conn),
		ContainerService: pb.NewContainerServiceClient(conn),
	}
	return &c
}
