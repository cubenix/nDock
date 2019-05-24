package controller

import (
	"html/template"

	"github.com/gauravgahlot/dockerdoodle/client/rpc"
	"github.com/gauravgahlot/dockerdoodle/types"
)

type container struct {
	hosts             *[]types.Host
	containerTemplate *template.Template
	client            rpc.ContainerServiceClient
}

func (c container) registerRoutes() {

}
