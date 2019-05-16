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
	for i := 0; i < len(hosts); i++ {
		h := vm.Host{
			Name:           hosts[i].Name,
			IP:             hosts[i].IP,
			ContainerCount: int(r.HostContainers[i].Containers[hosts[i].IP]),
			BGColor:        constants.BGClasses[i],
			ColorCode:      constants.TextClasses[i],
			BGColorCode:    constants.BGCodes[i],
		}
		res = append(res, h)
	}
	return &res
}
