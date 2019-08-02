package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/gauravgahlot/dockerdoodle/pkg/constants"
)

// StatsData represents the stats of a container
type StatsData struct {
	Index int32   `json:"index"`
	Usage float32 `json:"usage"`
}

var (
	// DoneCh is used to send a DONE signal
	DoneCh = make(chan struct{})

	// StatsCh holds the container stats
	StatsCh = make(chan StatsData)

	// DoneSignalSent some signal
	DoneSignalSent = true
)

// GetContainersCount returns the number of containers running on a host
func GetContainersCount(host string, all bool) (int, error) {
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
	return len(*c), nil
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
			sd := StatsData{Index: -1, Usage: 0.0}
			StatsCh <- sd
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

	sd := StatsData{Index: cIndex, Usage: cpuUsage(&st)}
	StatsCh <- sd
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

// StopContainer stops a container in running state
func StopContainer(ctx context.Context, host string, id string) error {
	cli, err := client.NewClientWithOpts(client.WithHost(constants.DockerAPIProtocol+host+constants.DockerAPIPort),
		client.WithVersion(constants.DockerAPIVersion))
	if err != nil {
		log.Fatal(err)
	}
	return cli.ContainerStop(ctx, id, nil)
}

// RemoveContainer removes a container in exited or created state
func RemoveContainer(ctx context.Context, host string, id string) error {
	cli, err := client.NewClientWithOpts(client.WithHost(constants.DockerAPIProtocol+host+constants.DockerAPIPort),
		client.WithVersion(constants.DockerAPIVersion))
	if err != nil {
		log.Fatal(err)
	}

	// TODO: accept options from user
	opts := types.ContainerRemoveOptions{
		Force:         false,
		RemoveLinks:   false,
		RemoveVolumes: true,
	}
	return cli.ContainerRemove(ctx, id, opts)
}
