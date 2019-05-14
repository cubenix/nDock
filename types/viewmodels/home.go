package viewmodels

// Host represents a Docker host
type Host struct {
	Name           string `json:"name"`
	IP             string `json:"ip"`
	ContainerCount int    `json:"containerCount"`
	BGColor        string `json:"bgColor"`
	ColorCode      string `json:"colorCode"`
	BGColorCode    string `json:"bgColorCode"`
}

// Home represents the view model for home page of WatchDock
type Home struct {
	Base
	Hosts []Host
}
