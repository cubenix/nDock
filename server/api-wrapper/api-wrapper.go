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
	StatsCh = make(chan map[int32]float32)

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

// GetContainers returns containers running on a host
func GetContainers(ctx context.Context, host string, quite bool, all bool) (*[]types.Container, error) {
	cli, err := client.NewClientWithOpts(client.WithHost(constants.DockerAPIProtocol+host+constants.DockerAPIPort),
		client.WithVersion(constants.DockerAPIVersion))
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	return getContainers(ctx, cli, quite, all)
}

func getContainers(ctx context.Context, cli *client.Client, quite bool, all bool) (*[]types.Container, error) {
	c, err := cli.ContainerList(ctx, types.ContainerListOptions{Quiet: quite, All: all})
	if err != nil {
		log.Fatal(err)
		return &[]types.Container{}, err
	}
	return &c, nil
}

// GetDockerStats returns CPU usage of a container
func GetDockerStats(ctx context.Context, host string, id string, cIndex int32) {
	cli, err := client.NewClientWithOpts(client.WithHost(constants.DockerAPIProtocol+host+constants.DockerAPIPort),
		client.WithVersion(constants.DockerAPIVersion))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-DoneCh:
			DoneSignalSent = true
			m := map[int32]float32{
				-1: 0.0,
			}
			StatsCh <- m
			return
		default:
			if !DoneSignalSent {
				getStats(ctx, cli, id, cIndex)
			}
		}
	}
}

func getStats(ctx context.Context, cli *client.Client, id string, cIndex int32) {
	s, e := cli.ContainerStats(ctx, id, false)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Body.Close()

	d, _ := ioutil.ReadAll(s.Body)
	var st types.Stats
	json.Unmarshal(d, &st)

	m := map[int32]float32{
		cIndex: cpuUsage(&st),
	}
	StatsCh <- m
}

func cpuUsage(stats *types.Stats) float32 {
	cpu := stats.CPUStats.CPUUsage.TotalUsage
	preCPU := stats.PreCPUStats.CPUUsage.TotalUsage
	systemCPU := stats.CPUStats.SystemUsage
	preSystemCPU := stats.PreCPUStats.SystemUsage
	cpuCount := len(stats.PreCPUStats.CPUUsage.PercpuUsage)

	// calculate the change for cpu usage of the container in between readings
	cpuDelta := cpu - preCPU

	// calculate the change for the entire system between readings
	systemDelta := systemCPU - preSystemCPU

	var usage float32
	if systemDelta > 0 && cpuDelta > 0 {
		usage = (float32(cpuDelta) / float32(systemDelta)) * float32(cpuCount) * float32(100)
	}
	return usage
}

// StartContainer starts a stopped or created container
func StartContainer(ctx context.Context, host string, id string) error {
	cli, err := client.NewClientWithOpts(client.WithHost(constants.DockerAPIProtocol+host+constants.DockerAPIPort),
		client.WithVersion(constants.DockerAPIVersion))
	if err != nil {
		log.Fatal(err)
	}

	return cli.ContainerStart(ctx, id, types.ContainerStartOptions{})
}

// StopContainer starts a stopped or created container
func StopContainer(ctx context.Context, host string, id string) error {
	cli, err := client.NewClientWithOpts(client.WithHost(constants.DockerAPIProtocol+host+constants.DockerAPIPort),
		client.WithVersion(constants.DockerAPIVersion))
	if err != nil {
		log.Fatal(err)
	}
	return cli.ContainerStop(ctx, id, nil)
}
