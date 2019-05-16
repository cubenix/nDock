package rpc

import (
	"github.com/gauravgahlot/dockerdoodle/pb"
	"google.golang.org/grpc"
)

// DockerServiceClient is the client for gRPC DockerService
type DockerServiceClient pb.DockerHostServiceClient

// ContainerServiceClient is the client for gRPC ContainerService
type ContainerServiceClient pb.ContainerServiceClient

// Client struct to hold all the service clients
type Client struct {
	DockerServiceClient    DockerServiceClient
	ContainerServiceClient ContainerServiceClient
}

// InitializeClients initializes client for respective service
func InitializeClients(conn *grpc.ClientConn) *Client {
	c := Client{
		DockerServiceClient:    pb.NewDockerHostServiceClient(conn),
		ContainerServiceClient: pb.NewContainerServiceClient(conn),
	}
	return &c
}
