package helpers

import (
	"context"
	"io"
	"log"

	convert "github.com/gauravgahlot/dockerdoodle/client/converters"
	vm "github.com/gauravgahlot/dockerdoodle/client/viewmodels"
	"github.com/gauravgahlot/dockerdoodle/pb"
	"github.com/gauravgahlot/dockerdoodle/types"
)

// GetContainersCount gets response from gRPC server
func GetContainersCount(c pb.DockerHostServiceClient, hosts *[]types.Host, all bool) (*[]vm.Host, error) {
	res, err := c.GetContainersCount(context.Background(), convert.ToGetContainersCountRequest(hosts, all))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return convert.ToHostsViewModel(res, *hosts), nil
}

// GetContainers returns
func GetContainers(c pb.DockerHostServiceClient, host string) (*[]vm.Container, error) {
	ctx := context.Background()
	res, err := c.GetContainers(ctx, convert.ToGetContainersRequest(host))
	var containers []vm.Container
	if err != nil {
		log.Fatal(err)
		return &containers, err
	}
	model, ids := convert.ToContainersViewModel(res)
	go streamStats(ctx, c, host, ids)
	return model, nil
}

func streamStats(ctx context.Context, c pb.DockerHostServiceClient, host string, ids *[]string) {
	stream, err := c.GetStats(ctx, &pb.GetStatsRequest{Host: host, ContainerIds: *ids})
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println("RES: ", res.Stats)
	}
}
