package helpers

import (
	"context"
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
	res, err := c.GetContainers(context.Background(), convert.ToGetContainersRequest(host))
	var containers []vm.Container
	if err != nil {
		log.Fatal(err)
		return &containers, err
	}
	return convert.ToContainersViewModel(res), nil
}
