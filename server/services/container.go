package services

import (
	"context"

	"github.com/gauravgahlot/watchdock/pb"
)

// ContainerService is a gRPC service to serve requests for Docker containers
type ContainerService struct{}

// GetContainer returns a container object for a container ID
func (s *ContainerService) GetContainer(ctx context.Context, req *pb.GetContainerRequest) (*pb.GetContainerResponse, error) {
	res := pb.GetContainerResponse{Container: &pb.Container{
		Id:   req.Id,
		Name: "container-name",
	}}
	return &res, nil
}
