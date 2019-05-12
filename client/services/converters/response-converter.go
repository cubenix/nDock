package converters

import (
	"github.com/gauravgahlot/watchdock/pb"
	"github.com/gauravgahlot/watchdock/types"
)

// ToGetContainersCountResponse return collection of `map[hostname]containers`
func ToGetContainersCountResponse(r *pb.GetContainersCountResponse, hosts []types.Host) *[]map[string]int {
	res := []map[string]int{}
	for i := 0; i < len(hosts); i++ {
		m := make(map[string]int)
		m[hosts[i].Name] = int(r.HostContainers[i].Containers[hosts[i].Name])
		res = append(res, m)
	}
	return &res
}
