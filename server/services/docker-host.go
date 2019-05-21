package services

import (
	"context"
	"io"
	"log"

	"github.com/gauravgahlot/dockerdoodle/pb"
	api "github.com/gauravgahlot/dockerdoodle/server/api-wrapper"
	convert "github.com/gauravgahlot/dockerdoodle/server/converter"
)

// DockerHostService is a gRPC service to serve requests for Docker Hosts
type DockerHostService struct{}

// GetContainersCount returns the number of containers running on each host
func (s *DockerHostService) GetContainersCount(ctx context.Context, req *pb.GetContainersCountRequest) (*pb.GetContainersCountResponse, error) {
	if !api.DoneSignalSent {
		api.DoneCh <- struct{}{}
		api.DoneSignalSent = true
	}

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

// GetContainers returns ID and name of each container running on a host
func (s *DockerHostService) GetContainers(ctx context.Context, req *pb.GetContainersRequest) (*pb.GetContainersResponse, error) {
	if !api.DoneSignalSent {
		api.DoneCh <- struct{}{}
		api.DoneSignalSent = true
	}
	containers, err := api.GetContainers(req.Host, false, false)
	if err != nil {
		log.Fatal(err)
	}

	return convert.ToGetContainersResponse(containers), nil
}

// GetStats sends containers stats to the client via a stream
func (s *DockerHostService) GetStats(req *pb.GetStatsRequest, stream pb.DockerHostService_GetStatsServer) error {
	api.DoneSignalSent = false
	ctx := context.Background()

	for cID, cIndex := range req.Containers {
		go api.GetDockerStats(ctx, req.Host, cID, cIndex)
	}

	for data := range api.StatsCh {
		err := stream.Send(&pb.GetStatsReponse{Stats: data})
		if err != nil {
			log.Fatal(err)
			return io.EOF
		}
	}
	return io.EOF
}

func sendDataOverStream(stream pb.DockerHostService_GetStatsServer) {
	for data := range api.StatsCh {
		err := stream.Send(&pb.GetStatsReponse{Stats: data})
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
