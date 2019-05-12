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
		HostContainers: []*pb.HostContainerCount{},
	}

	for _, host := range req.Hosts {
		c := pb.HostContainerCount{Containers: make(map[string]int32)}
		c.Containers[host] = 2
		res.HostContainers = append(res.HostContainers, &c)
	}
	return &res, nil
}
