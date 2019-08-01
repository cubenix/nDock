package svc

import (
	"context"
	"errors"
	"log"
	"strings"

	api "github.com/gauravgahlot/dockerdoodle/app/api-wrapper"
	vm "github.com/gauravgahlot/dockerdoodle/app/viewmodels"
	convert "github.com/gauravgahlot/dockerdoodle/pkg/converters"
)

// GetContainer get container details for a container ID
func GetContainer(ctx context.Context, host string, id string) (*vm.Container, error) {
	containers, err := api.GetContainers(ctx, host, true, true)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, c := range *containers {
		if strings.EqualFold(c.ID, id) {
			return convert.ToContainerViewModel(&c), nil
		}
	}
	return nil, errors.New("No container found with ID: " + id)
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
