package converters

import (
	"github.com/gauravgahlot/watchdock/constants"
	"github.com/gauravgahlot/watchdock/pb"
	"github.com/gauravgahlot/watchdock/types"
	vm "github.com/gauravgahlot/watchdock/types/viewmodels"
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
