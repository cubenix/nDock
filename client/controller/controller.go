package controller

import (
	"html/template"
	"net/http"

	"github.com/gauravgahlot/dockerdoodle/client/rpc"
	"github.com/gauravgahlot/dockerdoodle/types"
)

var (
	homeController home
)

// Startup registers all the HTTP request handlers
func Startup(templates map[string]*template.Template, client *rpc.Client, hosts *[]types.Host) {
	homeController.homeTemplate = templates["home.html"]
	homeController.hosts = hosts
	homeController.client = client.DockerServiceClient
	homeController.registerRoutes()

	http.Handle("/js/", http.FileServer(http.Dir("client/public")))
	http.Handle("/vendor/", http.FileServer(http.Dir("client/public")))
	http.Handle("/css/", http.FileServer(http.Dir("client/public")))
}
