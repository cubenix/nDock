package apiwrapper

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/gauravgahlot/dockerdoodle/constants"
)

var (
	apiClient *client.Client
	doneCh    = make(chan struct{})
)

// InitializeClient initializes a new instance of the API Client
func InitializeClient(host string) {
	cli, err := client.NewClientWithOpts(client.WithHost(constants.DockerAPIProtocol+host+constants.DockerAPIPort), client.WithVersion(constants.DockerAPIVersion))
	if err != nil {
		log.Fatal(err)
	}
	apiClient = cli
}

// GetContainersCount returns the number of containers running on a host
func GetContainersCount(host string, all bool) (int32, error) {
	c, err := GetContainers(context.Background(), true, all)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	defer apiClient.Close()
	return int32(len(*c)), nil
}

// GetContainers returns containers running on a host
func GetContainers(ctx context.Context, quite bool, all bool) (*[]types.Container, error) {
	c, err := apiClient.ContainerList(ctx, types.ContainerListOptions{Quiet: quite, All: all})
	if err != nil {
		log.Fatal(err)
		return &[]types.Container{}, err
	}
	return &c, nil
}

func containerStats(ctx context.Context, cli *client.Client, host string) {
	c, _ := cli.ContainerList(ctx, types.ContainerListOptions{Quiet: true})
	defer log.Println("closing containerStats")

	go getDockerStats(ctx, cli, c[0].ID)
	time.Sleep(5 * time.Second)
	doneCh <- struct{}{}
}

func getDockerStats(ctx context.Context, cli *client.Client, id string) {
	defer log.Println("closing getDockerStats")
	for {
		select {
		case <-doneCh:
			log.Println("received DONE signal")
			return
		default:
			getStats(ctx, cli, id)
		}
	}
}

func getStats(ctx context.Context, cli *client.Client, id string) {
	s, e := cli.ContainerStats(ctx, id, false)
	if e == io.EOF {
		log.Println("EoF")
	}
	if e != nil {
		log.Fatal(e)
	}
	defer func() {
		log.Println("closing Body")
		s.Body.Close()
	}()

	d, _ := ioutil.ReadAll(s.Body)
	var st types.Stats
	json.Unmarshal(d, &st)
	log.Println()
	log.Println("CPU Usage", st.CPUStats.CPUUsage)
}
