package helpers

import (
	"context"
	"encoding/json"
	"io"
	"log"

	convert "github.com/gauravgahlot/dockerdoodle/client/converters"
	vm "github.com/gauravgahlot/dockerdoodle/client/viewmodels"
	"github.com/gauravgahlot/dockerdoodle/client/ws"
	"github.com/gauravgahlot/dockerdoodle/pkg/pb"
	"github.com/gauravgahlot/dockerdoodle/pkg/types"
)

// Hub is a hub of clients and channels
var Hub *ws.Hub

// GetContainersCount gets response from gRPC server
func GetContainersCount(c pb.DockerHostServiceClient, hosts *[]types.Host, all bool) (*[]vm.Host, error) {
	res, err := c.GetContainersCount(context.Background(), convert.ToGetContainersCountRequest(hosts, all))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return convert.ToHostsViewModel(res, *hosts), nil
}

// GetContainers returns a pointer to collection of container view model
func GetContainers(ctx context.Context, c pb.DockerHostServiceClient, host string, stats bool) (*[]vm.Container, *[]vm.Container, error) {
	res, err := c.GetContainers(ctx, convert.ToGetContainersRequest(host))
	var containers []vm.Container
	if err != nil {
		log.Fatal(err)
		return nil, &containers, err
	}
	all, running, req := convert.ToContainersViewModelAndGetStatsRequest(res, host)
	if stats {
		go streamStats(ctx, c, req)
	}
	return all, running, nil
}

func streamStats(ctx context.Context, c pb.DockerHostServiceClient, req *pb.GetStatsRequest) {
	type streamData struct {
		Index int32   `json:"index"`
		Usage float32 `json:"usage"`
	}

	stream, err := c.GetStats(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		} else if err != nil {
			log.Fatal("received ERROR: ", err)
			return
		}

		var data streamData
		for i, d := range res.Stats {
			data.Index = i
			data.Usage = d
		}

		if data, er := json.Marshal(data); er == nil {
			Hub.Broadcast <- data
		}
	}
}
