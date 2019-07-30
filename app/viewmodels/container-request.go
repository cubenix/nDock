package viewmodels

// ContainerOperationRequest represents the request for container operations:
// START, STOP, REMOVE
type ContainerOperationRequest struct {
	Host string `json:"host"`
	ID   string `json:"id"`
}
