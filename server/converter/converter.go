package converter

import (
	"github.com/docker/docker/api/types"
	"github.com/gauravgahlot/dockerdoodle/pkg/pb"
)

// ToGetContainersResponse returns response object for GetContainers call
func ToGetContainersResponse(containers *[]types.Container) *pb.GetContainersResponse {
	res := pb.GetContainersResponse{Containers: []*pb.Container{}}
	for _, c := range *containers {
		res.Containers = append(res.Containers, getContainer(&c))
	}
	return &res
}

// ToGetContainerResponse returns response object for GetContainer call
func ToGetContainerResponse(c *types.Container) *pb.GetContainerResponse {
	return &pb.GetContainerResponse{Container: getContainer(c)}
}

func getContainer(c *types.Container) *pb.Container {
	pc := pb.Container{
		Id:      c.ID,
		Name:    c.Names[0][1:],
		Image:   c.Image,
		Command: c.Command,
		Created: c.Created,
		State:   c.State,
		Status:  c.Status,
		Ports:   []*pb.Port{},
		Mounts:  []*pb.MountPoint{},
	}

	for _, p := range c.Ports {
		port := pb.Port{
			IP:          p.IP,
			PrivatePort: int32(p.PrivatePort),
			PublicPort:  int32(p.PublicPort),
			Type:        p.Type,
		}
		pc.Ports = append(pc.Ports, &port)
	}

	for _, m := range c.Mounts {
		mp := pb.MountPoint{
			Type:        string(m.Type),
			Name:        m.Name,
			Source:      m.Source,
			Destination: m.Destination,
			Mode:        m.Mode,
			RW:          m.RW,
		}
		pc.Mounts = append(pc.Mounts, &mp)
	}
	return &pc
}
