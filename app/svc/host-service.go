package svc

import (
	"context"
	"log"

	api "github.com/gauravgahlot/dockerdoodle/app/api-wrapper"
	"github.com/gauravgahlot/dockerdoodle/pkg/types"
)

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
