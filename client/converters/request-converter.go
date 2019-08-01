package converters

import (
	"github.com/gauravgahlot/dockerdoodle/pkg/pb"
	"github.com/gauravgahlot/dockerdoodle/pkg/types"
)

// ToGetContainersCountRequest returns GetContainersCountRequestObject
func ToGetContainersCountRequest(hosts *[]types.Host, all bool) *pb.GetContainersCountRequest {
	req := pb.GetContainersCountRequest{Hosts: []string{}, All: all}
	for _, h := range *hosts {
		req.Hosts = append(req.Hosts, h.IP)
	}
	return &req
}

// ToGetContainersRequest return ToGetContainersRequestObject
func ToGetContainersRequest(host string) *pb.GetContainersRequest {
	return &pb.GetContainersRequest{Host: host}
}
