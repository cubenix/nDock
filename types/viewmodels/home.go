package viewmodels

// Host represents a Docker host
type Host struct {
	Name           string
	IP             string
	ContainerCount int
	BGColor        string
	ColorCode      string
	BGColorCode    string
}

// Home represents the view model for home page of WatchDock
type Home struct {
	Base
	Hosts []Host
}
