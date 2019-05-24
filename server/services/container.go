package services

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/gauravgahlot/dockerdoodle/pb"
	api "github.com/gauravgahlot/dockerdoodle/server/api-wrapper"
	"github.com/gauravgahlot/dockerdoodle/server/converter"
)

// ContainerService is a gRPC service to serve requests for Docker containers
type ContainerService struct{}

// GetContainer returns a container object for a container ID
func (s *ContainerService) GetContainer(ctx context.Context, req *pb.GetContainerRequest) (*pb.GetContainerResponse, error) {
	containers, err := api.GetContainers(ctx, req.Host, true, true)

	if err != nil {
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	for _, c := range *containers {
		if strings.EqualFold(c.ID, req.ID) {
			return converter.ToGetContainerResponse(&c), nil
		}
	}
	return nil, errors.New("No container found with ID: " + req.ID)
}

// StartContainer starts a stopped or created container, if there is no error
func (s *ContainerService) StartContainer(ctx context.Context, req *pb.StartContainerRequest) (*pb.ErrorStatus, error) {
	err := api.StartContainer(ctx, req.Host, req.ID)
	if err != nil {
		return &pb.ErrorStatus{Message: "Failed to start the container"}, err
	}
	return nil, nil
}

// StopContainer starts a stopped or created container, if there is no error
func (s *ContainerService) StopContainer(ctx context.Context, req *pb.StopContainerRequest) (*pb.ErrorStatus, error) {
	err := api.StopContainer(ctx, req.Host, req.ID)
	if err != nil {
		return &pb.ErrorStatus{Message: "Failed to stop the container"}, err
	}
	return nil, nil
}
