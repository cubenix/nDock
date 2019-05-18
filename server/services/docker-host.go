package services

import (
	"context"
	"log"

	"github.com/gauravgahlot/dockerdoodle/pb"
	api "github.com/gauravgahlot/dockerdoodle/server/api-wrapper"
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
		api.InitializeClient(host)
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

// GetContainers returns ID and name of each container running on a host
func (s *DockerHostService) GetContainers(ctx context.Context, req *pb.GetContainersRequest) (*pb.GetContainersResponse, error) {
	api.InitializeClient(req.Host)
	containers, err := api.GetContainers(context.Background(), false, false)
	if err != nil {
		log.Fatal(err)
	}
	res := pb.GetContainersResponse{Containers: []*pb.Container{}}
	for _, c := range *containers {
		res.Containers = append(res.Containers, &pb.Container{Id: c.ID, Name: c.Names[0][1:]})
	}
	return &res, nil
}

// GetStats sends containers stats to the client via a stream
func (s *DockerHostService) GetStats(req *pb.GetStatsRequest, stream pb.DockerHostService_GetStatsServer) error {
	stream.Send(&pb.GetStatsReponse{})
	return nil
}
