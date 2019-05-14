package converters

import (
	"github.com/gauravgahlot/watchdock/pb"
	"github.com/gauravgahlot/watchdock/types"
)

// ToGetContainersCountRequest returns GetContainersCountRequestObject
func ToGetContainersCountRequest(hosts *[]types.Host, all bool) *pb.GetContainersCountRequest {
	req := pb.GetContainersCountRequest{Hosts: []string{}, All: all}
	for _, h := range *hosts {
		req.Hosts = append(req.Hosts, h.Name)
	}
	return &req
}
