package svc

import (
	"context"
	"log"

	vm "github.com/gauravgahlot/dockerdoodle/app/viewmodels"
	convert "github.com/gauravgahlot/dockerdoodle/pkg/converters"
	"github.com/gauravgahlot/dockerdoodle/pkg/pb"
	api "github.com/gauravgahlot/dockerdoodle/app/api-wrapper"
)

// GetContainer get container details for a container ID
func GetContainer(ctx context.Context, host string, id string) (*vm.Container, error) {
	err := c.GetContainer(ctx, &pb.GetContainerRequest{ID: id, Host: host})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return convert.ToContainerViewModel(res.Container), nil
}

// StartContainer starts a stopped or created container, if there is no error
func StartContainer(ctx context.Context, host string, id string) error {
	err := api.StartContainer(ctx, host, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// StopContainer stops a running container, if there is no error
func StopContainer(ctx context.Context, host string, id string) error {
	err := api.StopContainer(ctx, host, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// RemoveContainer removes a container in exited or created state, if there is no error
func RemoveContainer(ctx context.Context, host string, id string) error {
	err := api.RemoveContainer(ctx, host, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
