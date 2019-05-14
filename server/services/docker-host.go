package services

import (
	"context"
	"log"

	"github.com/gauravgahlot/watchdock/pb"
	api "github.com/gauravgahlot/watchdock/server/api-wrapper"
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

		count, err := api.GetContainersCount(host, req.All)
		if err != nil {
			log.Fatal(err)
			c.Containers[host] = -1
		} else {
			c.Containers[host] = count
		}
		res.HostContainers = append(res.HostContainers, &c)
	}
	return &res, nil
}
