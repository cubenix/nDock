package helpers

import (
	"context"
	"log"

	convert "github.com/gauravgahlot/watchdock/client/services/converters"
	"github.com/gauravgahlot/watchdock/pb"
	"github.com/gauravgahlot/watchdock/types"
)

// GetContainersCount gets response from gRPC server
func GetContainersCount(c pb.DockerHostServiceClient, hosts *[]types.Host) (*[]types.Host, error) {
	res, err := c.GetContainersCount(context.Background(), convert.ToGetContainersCountRequest(hosts))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return convert.ToHostsViewModel(res, *hosts), nil
}
