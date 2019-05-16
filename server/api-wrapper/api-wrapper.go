package apiwrapper

import (
	"context"
	"io"
	"log"

	"github.com/docker/docker/api/types"

	"github.com/docker/docker/client"
	"github.com/gauravgahlot/dockerdoodle/constants"
)

// GetContainersCount returns the number of containers running on a host
func GetContainersCount(host string, all bool) (int32, error) {
	cli, cliErr := client.NewClientWithOpts(client.WithHost(constants.DockerAPIProtocol+host+constants.DockerAPIPort), client.WithVersion(constants.DockerAPIVersion))
	if cliErr != nil {
		log.Fatal(cliErr)
		return -1, cliErr
	}

	go containerStats(host)

	c, err := cli.ContainerList(context.Background(), types.ContainerListOptions{Quiet: true, All: all})
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	return int32(len(c)), nil
}

func containerStats(host string) {
	cli, _ := client.NewClientWithOpts(client.WithHost(constants.DockerAPIProtocol+host+constants.DockerAPIPort), client.WithVersion(constants.DockerAPIVersion))
	containers, _ := cli.ContainerList(context.Background(), types.ContainerListOptions{Quiet: true})

	go streamStats(cli, containers[0].ID)
}

func streamStats(cli *client.Client, id string) {
	for {
		s, e := cli.ContainerStats(context.Background(), id, true)
		if e == io.EOF {
			log.Println("EoF")
		}
		// if e != nil {
		// 	break
		// }
		log.Println("OS Type: ", s.OSType)
		var data []byte
		s.Body.Read(data)
		s.Body.Close()
		log.Println(string(data))
	}
}
