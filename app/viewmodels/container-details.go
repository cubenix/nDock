package viewmodels

// ContainerDetails is the view model for container details page
type ContainerDetails struct {
	Base
	Container    Container
	SelectedHost string
}
