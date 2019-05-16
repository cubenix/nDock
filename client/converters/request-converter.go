package converters

import (
	"github.com/gauravgahlot/dockerdoodle/pb"
	"github.com/gauravgahlot/dockerdoodle/types"
)

// ToGetContainersCountRequest returns GetContainersCountRequestObject
func ToGetContainersCountRequest(hosts *[]types.Host, all bool) *pb.GetContainersCountRequest {
	req := pb.GetContainersCountRequest{Hosts: []string{}, All: all}
	for _, h := range *hosts {
		req.Hosts = append(req.Hosts, h.IP)
	}
	return &req
}
