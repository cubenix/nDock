package viewmodels

// Container represents a Docker container
type Container struct {
	ID        string
	Name      string
	Image     string
	Command   string
	Created   int64
	State     string
	Status    string
	Ports     []Port
	Mounts    []Mount
	ColorCode string
}

// Port holds all the details of a port mapping for a container
type Port struct {
	IP          string
	PrivatePort int32
	PublicPort  int32
	Type        string
}

// Mount holds all the details of a mount point in a container
type Mount struct {
	Type        string
	Name        string
	Source      string
	Destination string
	Mode        string
	RW          bool
}
