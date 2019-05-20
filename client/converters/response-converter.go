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

// ToContainersViewModel returns pointers to collection of Container view model and GetStatsRequest
func ToContainersViewModel(r *pb.GetContainersResponse, host string) (*[]vm.Container, *pb.GetStatsRequest) {
	res := []vm.Container{}
	req := pb.GetStatsRequest{Host: host, Containers: map[string]int32{}}

	for i, c := range r.Containers {
		res = append(res, vm.Container{
			ID:        c.Id,
			Name:      c.Name,
			ColorCode: constants.BGCodes[i],
		})
		req.Containers[c.Id] = int32(i)
	}
	return &res, &req
}
