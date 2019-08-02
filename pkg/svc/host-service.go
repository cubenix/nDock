package svc

import (
	"context"
	"encoding/json"
	"log"

	vm "github.com/gauravgahlot/dockerdoodle/app/viewmodels"
	"github.com/gauravgahlot/dockerdoodle/app/ws"
	"github.com/gauravgahlot/dockerdoodle/pkg/api"
	cnv "github.com/gauravgahlot/dockerdoodle/pkg/converters"
	"github.com/gauravgahlot/dockerdoodle/pkg/types"
)

// Hub is a hub of clients and channels
var Hub *ws.Hub

// GetContainersCount returns the number of containers running on each host
func GetContainersCount(ctx context.Context, hosts *[]types.Host, all bool) (*map[string]int, error) {
	if !api.DoneSignalSent {
		api.DoneCh <- struct{}{}
		api.DoneSignalSent = true
	}

	res := make(map[string]int)
	for _, host := range *hosts {
		count, err := api.GetContainersCount(host.IP, all)
		if err != nil {
			log.Fatal(err)
			res[host.IP] = -1
		} else {
			res[host.IP] = count
		}
	}
	return &res, nil
}

// GetContainers returns a pointer to collection of container view model
func GetContainers(ctx context.Context, host string, stats bool) (*[]vm.Container, *[]vm.Container, error) {
	res, err := api.GetContainers(ctx, host, false, true)
	var containers []vm.Container
	if err != nil {
		log.Fatal(err)
		return nil, &containers, err
	}
	all, running, req := cnv.ToContainersViewModelAndGetStatsRequest(res, host)
	if stats {
		api.DoneSignalSent = false
		go streamStats(ctx, host, req)
	}
	return all, running, nil
}

func streamStats(ctx context.Context, host string, req *map[string]int32) {
	for cID, cIndex := range *req {
		go api.GetDockerStats(ctx, host, cID, cIndex)
	}

	for data := range api.StatsCh {
		if data.Index == -1 {
			break
		}
		if data, er := json.Marshal(data); er == nil {
			Hub.Broadcast <- data
		}
	}
}
