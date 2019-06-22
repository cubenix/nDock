package main

import (
	"flag"

	"github.com/gauravgahlot/dockerdoodle/client"
	"github.com/gauravgahlot/dockerdoodle/constants"
	"github.com/gauravgahlot/dockerdoodle/server"
)

var (
	useLocal       = flag.Bool("L", false, "use localhost as the only Docker Host")
	serverEndpoint = flag.String("s", constants.LocalIP, "endpoint of the GRPC server")
)

func main() {
	flag.Parse()

	// initialize gRPC server
	go server.Initialize()

	// initialize gRPC client
	client.Initialize(*useLocal, *serverEndpoint)
}
