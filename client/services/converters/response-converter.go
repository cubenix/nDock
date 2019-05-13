package converters

import (
	"github.com/gauravgahlot/watchdock/constants"
	"github.com/gauravgahlot/watchdock/pb"
	"github.com/gauravgahlot/watchdock/types"
)

// ToHostsViewModel returns a collection of Hosts
func ToHostsViewModel(r *pb.GetContainersCountResponse, hosts []types.Host) *[]types.Host {
	res := []types.Host{}
	for i := 0; i < len(hosts); i++ {
		hosts[i].ContainerCount = int(r.HostContainers[i].Containers[hosts[i].Name])
		hosts[i].BGColor = constants.BGClasses[i]
		res = append(res, hosts[i])
	}
	return &res
}
