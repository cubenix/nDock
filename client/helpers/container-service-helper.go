package helpers

import (
	"context"
	"log"

	convert "github.com/gauravgahlot/dockerdoodle/client/converters"
	vm "github.com/gauravgahlot/dockerdoodle/client/viewmodels"
	"github.com/gauravgahlot/dockerdoodle/pb"
)

// GetContainer get container details for a container ID
func GetContainer(ctx context.Context, c pb.ContainerServiceClient, host string, id string) (*vm.Container, error) {
	res, err := c.GetContainer(ctx, &pb.GetContainerRequest{ID: id, Host: host})
	if err != nil {
		log.Fatal(err)
	}
	return convert.ToContainerViewModel(res.Container), nil
}
