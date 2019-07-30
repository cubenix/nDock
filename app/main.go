package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gauravgahlot/dockerdoodle/app/controller"
	"github.com/gauravgahlot/dockerdoodle/pkg/constants"
	"github.com/gauravgahlot/dockerdoodle/pkg/helpers"
	"github.com/gauravgahlot/dockerdoodle/pkg/types"
)

var (
	useLocal       = flag.Bool("L", false, "use localhost as the only Docker Host")
	serverEndpoint = flag.String("s", constants.LocalIP, "endpoint of the GRPC server")
)

func main() {
	flag.Parse()
	templates := helpers.PopulateTemplates()

	var config *types.Config
	if *useLocal {
		config = helpers.ConfigForLocalEnv()
	} else {
		config = helpers.ReadConfiguration()
	}

	controller.Startup(templates, &config.Hosts)
	server := http.Server{
		Addr: constants.ApplicationPort,
	}
	log.Println("DockerDoodle listening at port", constants.ApplicationPort)
	log.Fatal(server.ListenAndServe())
}
