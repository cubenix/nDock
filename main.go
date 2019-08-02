package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gauravgahlot/dockerdoodle/app/controller"
	"github.com/gauravgahlot/dockerdoodle/pkg/constants"
	"github.com/gauravgahlot/dockerdoodle/pkg/svc"
	"github.com/gauravgahlot/dockerdoodle/pkg/types"
)

var (
	useLocal = flag.Bool("L", false, "use localhost as the only Docker Host")
)

func main() {
	flag.Parse()
	templates := svc.PopulateTemplates()

	var config *types.Config
	if *useLocal {
		config = svc.ConfigForLocalEnv()
	} else {
		config = svc.ReadConfiguration()
	}

	controller.Startup(templates, &config.Hosts)
	server := http.Server{
		Addr: constants.ApplicationPort,
	}
	log.Println("DockerDoodle listening at port", constants.ApplicationPort)
	log.Fatal(server.ListenAndServe())
}
