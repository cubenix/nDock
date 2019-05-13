package types

// Host represents a Docker host
type Host struct {
	Name           string
	IP             string
	ContainerCount int
	BGColor        string
}
