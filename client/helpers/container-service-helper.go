package helpers

import (
	"context"
	"errors"
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
		return nil, err
	}
	return convert.ToContainerViewModel(res.Container), nil
}

// StartContainer starts a stopped or created container
func StartContainer(ctx context.Context, c pb.ContainerServiceClient, host string, id string) error {
	res, err := c.StartContainer(ctx, &pb.StartContainerRequest{ID: id, Host: host})
	if err != nil {
		log.Fatal(err)
		return errors.New(res.Message)
	}
	return nil
}

// StopContainer starts a stopped or created container
func StopContainer(ctx context.Context, c pb.ContainerServiceClient, host string, id string) error {
	res, err := c.StopContainer(ctx, &pb.StopContainerRequest{ID: id, Host: host})
	if err != nil {
		log.Fatal(err)
		return errors.New(res.Message)
	}
	return nil
}
