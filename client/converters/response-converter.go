package converters

import (
	vm "github.com/gauravgahlot/dockerdoodle/client/viewmodels"
	"github.com/gauravgahlot/dockerdoodle/constants"
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

// ToContainersViewModel returns a collection of Container view model and container IDs
func ToContainersViewModel(r *pb.GetContainersResponse) (*[]vm.Container, *[]string) {
	res := []vm.Container{}
	ids := []string{}
	for i, c := range r.Containers {
		res = append(res, vm.Container{
			ID:        c.Id,
			Name:      c.Name,
			ColorCode: constants.BGCodes[i],
		})
		ids = append(ids, c.Id)
	}
	return &res, &ids
}
