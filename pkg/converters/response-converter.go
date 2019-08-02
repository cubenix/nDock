package converters

import (
	dt "github.com/docker/docker/api/types"
	vm "github.com/gauravgahlot/dockerdoodle/app/viewmodels"
	"github.com/gauravgahlot/dockerdoodle/pkg/constants"
	"github.com/gauravgahlot/dockerdoodle/pkg/types"
)

// ToHostsViewModel returns a collection of Host view model
func ToHostsViewModel(r map[string]int, hosts []types.Host) *[]vm.Host {
	res := []vm.Host{}
	for _, host := range hosts {
		h := vm.Host{
			Name:           host.Name,
			IP:             host.IP,
			ContainerCount: int(r[host.IP]),
		}
		res = append(res, h)
	}
	return &res
}

// ToContainersViewModelAndGetStatsRequest returns pointers to collection of Container view model and GetStatsRequest
func ToContainersViewModelAndGetStatsRequest(res *[]dt.Container, host string) (*[]vm.Container, *[]vm.Container, *map[string]int32) {
	all := []vm.Container{}
	running := []vm.Container{}
	req := map[string]int32{}

	for i, c := range *res {
		ct := ToContainerViewModel(&c)
		ct.ColorCode = constants.BGCodes[i]
		all = append(all, *ct)

		if c.State == constants.ContainerRunning {
			running = append(running, *ct)
		}
	}

	for i, c := range running {
		req[c.ID] = int32(i)
	}
	return &all, &running, &req
}

// ToContainerViewModel returns a struct of Container view model
func ToContainerViewModel(c *dt.Container) *vm.Container {
	return &vm.Container{
		ID:      c.ID,
		Name:    c.Names[0][1:],
		Image:   c.Image,
		Command: c.Command,
		Created: c.Created,
		State:   c.State,
		Status:  c.Status,
		Ports:   *getPorts(c.Ports),
		Mounts:  *getMounts(c.Mounts),
	}
}

func getPorts(ports []dt.Port) *[]vm.Port {
	ps := []vm.Port{}
	for _, p := range ports {
		ps = append(ps, vm.Port{
			IP:          p.IP,
			Type:        p.Type,
			PrivatePort: int32(p.PrivatePort),
			PublicPort:  int32(p.PublicPort),
		})
	}
	return &ps
}

func getMounts(mounts []dt.MountPoint) *[]vm.Mount {
	ms := []vm.Mount{}
	for _, m := range mounts {
		ms = append(ms, vm.Mount{
			Type:        string(m.Type),
			Name:        m.Name,
			Source:      m.Source,
			Destination: m.Destination,
			Mode:        m.Mode,
			RW:          m.RW,
		})
	}
	return &ms
}
