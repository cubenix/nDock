package helpers

import (
	"context"
	"log"

	convert "github.com/gauravgahlot/watchdock/client/services/converters"
	"github.com/gauravgahlot/watchdock/pb"
	"github.com/gauravgahlot/watchdock/types"
	vm "github.com/gauravgahlot/watchdock/types/viewmodels"
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
