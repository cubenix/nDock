package services

import (
	"context"

	"github.com/gauravgahlot/watchdock/pb"
)

// DockerHostService is a gRPC service to serve requests for Docker Hosts
type DockerHostService struct{}

// GetContainersCount returns the number of containers running on each host
func (s *DockerHostService) GetContainersCount(ctx context.Context, req *pb.GetContainersCountRequest) (*pb.GetContainersCountResponse, error) {
	res := pb.GetContainersCountResponse{
		HostContainers: []*pb.GetContainersCountResponse_HostContainerMapField{
			&pb.GetContainersCountResponse_HostContainerMapField{
				Host:  "hostname",
				Count: 2,
			},
		},
	}
	return &res, nil
}
