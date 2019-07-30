package controller

import (
	"html/template"
	"net/http"

	"github.com/gauravgahlot/dockerdoodle/pkg/types"
)

var (
	homeController      home
	hostController      host
	containerController container
)

// Startup registers all the HTTP request handlers
func Startup(templates map[string]*template.Template, hosts *[]types.Host) {
	homeController.homeTemplate = templates["home.html"]
	homeController.hosts = hosts
	homeController.registerRoutes()

	hostController.hostTemplate = templates["host.html"]
	hostController.hostContainerTemplate = templates["host-containers.html"]
	hostController.hosts = hosts
	hostController.registerRoutes()

	containerController.containerTemplate = templates["container.html"]
	containerController.hosts = hosts
	containerController.registerRoutes()

	http.Handle("/js/", http.FileServer(http.Dir("client/public")))
	http.Handle("/img/", http.FileServer(http.Dir("client/public")))
	http.Handle("/vendor/", http.FileServer(http.Dir("client/public")))
	http.Handle("/css/", http.FileServer(http.Dir("client/public")))
}
