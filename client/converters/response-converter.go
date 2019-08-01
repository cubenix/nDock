package converters

import (
	vm "github.com/gauravgahlot/dockerdoodle/client/viewmodels"
	"github.com/gauravgahlot/dockerdoodle/pkg/constants"
	"github.com/gauravgahlot/dockerdoodle/pkg/pb"
	"github.com/gauravgahlot/dockerdoodle/pkg/types"
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

// ToContainersViewModelAndGetStatsRequest returns pointers to collection of Container view model and GetStatsRequest
func ToContainersViewModelAndGetStatsRequest(r *pb.GetContainersResponse, host string) (*[]vm.Container, *[]vm.Container, *pb.GetStatsRequest) {
	all := []vm.Container{}
	running := []vm.Container{}
	req := pb.GetStatsRequest{Host: host, Containers: map[string]int32{}}

	for i, c := range r.Containers {
		ct := ToContainerViewModel(c)
		ct.ColorCode = constants.BGCodes[i]
		all = append(all, *ct)

		if c.State == constants.ContainerRunning {
			running = append(running, *ct)
		}
	}

	for i, c := range running {
		req.Containers[c.ID] = int32(i)
	}
	return &all, &running, &req
}

// ToContainerViewModel returns a struct of Container view model
func ToContainerViewModel(c *pb.Container) *vm.Container {
	return &vm.Container{
		ID:      c.Id,
		Name:    c.Name,
		Image:   c.Image,
		Command: c.Command,
		Created: c.Created,
		State:   c.State,
		Status:  c.Status,
		Ports:   *getPorts(c.Ports),
		Mounts:  *getMounts(c.Mounts),
	}
}

func getPorts(ports []*pb.Port) *[]vm.Port {
	ps := []vm.Port{}
	for _, p := range ports {
		ps = append(ps, vm.Port{
			IP:          p.IP,
			Type:        p.Type,
			PrivatePort: p.PrivatePort,
			PublicPort:  p.PublicPort,
		})
	}
	return &ps
}

func getMounts(mounts []*pb.MountPoint) *[]vm.Mount {
	ms := []vm.Mount{}
	for _, m := range mounts {
		ms = append(ms, vm.Mount{
			Type:        m.Type,
			Name:        m.Name,
			Source:      m.Source,
			Destination: m.Destination,
			Mode:        m.Mode,
			RW:          m.RW,
		})
	}
	return &ms
}
