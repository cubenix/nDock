package viewmodels

// HostContainers represents the view model for home page of WatchDock
type HostContainers struct {
	Base
	Hosts             []Host
	SelectedHost      string
	AllContainers     []Container
	RunningContainers []Container
}
