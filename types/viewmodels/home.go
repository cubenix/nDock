package viewmodels

import (
	"github.com/gauravgahlot/watchdock/types"
)

// Home represents the view model for home page of WatchDock
type Home struct {
	Base
	Hosts []types.Host
}
