package converters

import (
	vm "github.com/gauravgahlot/dockerdoodle/client/viewmodels"
	"github.com/gauravgahlot/dockerdoodle/pb"
	"github.com/gauravgahlot/dockerdoodle/types"
)

// ToHostsViewModel returns a collection of Host view model
func ToHostsViewModel(r *pb.GetContainersCountResponse, hosts []types.Host) *[]vm.Host {
	res := []vm.Host{}
	for i, host := range hosts {
		h := vm.Host{
			Name:           host.Name,
			IP:             host.IP,
			ContainerCount: int(r.HostContainers[i].Containers[host.IP]),
		}
		res = append(res, h)
	}
	return &res
}
