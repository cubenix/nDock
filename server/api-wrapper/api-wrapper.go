package apiwrapper

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"

	"github.com/docker/docker/client"
	"github.com/gauravgahlot/watchdock/constants"
)

// GetContainersCount returns the number of containers running on a host
func GetContainersCount(host string, all bool) (int32, error) {
	cli, cliErr := client.NewClientWithOpts(client.WithHost(constants.DockerAPIProtocol+host+constants.DockerAPIPort), client.WithVersion(constants.DockerAPIVersion))
	if cliErr != nil {
		log.Fatal(cliErr)
		return -1, cliErr
	}

	c, err := cli.ContainerList(context.Background(), types.ContainerListOptions{Quiet: true, All: all})
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	return int32(len(c)), nil
}
