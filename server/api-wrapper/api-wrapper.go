package apiwrapper

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/gauravgahlot/dockerdoodle/constants"
)

var (
	// DoneCh is used to send a DONE signal
	DoneCh = make(chan struct{})

	// StatsCh holds the container stats
	StatsCh = make(chan map[string]float32)

	// DoneSignalSent some signal
	DoneSignalSent = true
)

// GetContainersCount returns the number of containers running on a host
func GetContainersCount(host string, all bool) (int32, error) {
	cli, err := client.NewClientWithOpts(client.WithHost(constants.DockerAPIProtocol+host+constants.DockerAPIPort), client.WithVersion(constants.DockerAPIVersion))
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	c, err := getContainers(context.Background(), cli, true, all)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	return int32(len(*c)), nil
}

func getContainers(ctx context.Context, cli *client.Client, quite bool, all bool) (*[]types.Container, error) {
	c, err := cli.ContainerList(ctx, types.ContainerListOptions{Quiet: quite, All: all})
	if err != nil {
		log.Fatal(err)
		return &[]types.Container{}, err
	}
	return &c, nil
}

// GetContainers returns containers running on a host
func GetContainers(host string, quite bool, all bool) (*[]types.Container, error) {
	cli, err := client.NewClientWithOpts(client.WithHost(constants.DockerAPIProtocol+host+constants.DockerAPIPort),
		client.WithVersion(constants.DockerAPIVersion))
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	return getContainers(context.Background(), cli, quite, all)
}

// GetDockerStats returns CPU usage of a container
func GetDockerStats(ctx context.Context, host string, id string) {
	cli, err := client.NewClientWithOpts(client.WithHost(constants.DockerAPIProtocol+host+constants.DockerAPIPort),
		client.WithVersion(constants.DockerAPIVersion))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-DoneCh:
			DoneSignalSent = true
			cli.Close()
			return
		default:
			if !DoneSignalSent {
				getStats(ctx, cli, id)
			}
		}
	}
}

func getStats(ctx context.Context, cli *client.Client, id string) {
	s, e := cli.ContainerStats(ctx, id, false)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Body.Close()

	d, _ := ioutil.ReadAll(s.Body)
	var st types.Stats
	json.Unmarshal(d, &st)

	// TODO: calculate CPU utilization and push to channel
	m := make(map[string]float32)
	m[id] = float32(st.CPUStats.CPUUsage.TotalUsage)
	StatsCh <- m
}
