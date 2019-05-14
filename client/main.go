package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gauravgahlot/watchdock/client/rpc"
	"github.com/gauravgahlot/watchdock/client/services"
	"github.com/gauravgahlot/watchdock/constants"

	"google.golang.org/grpc"
)

var handler *services.Handler

type base struct {
	Title string
}

func init() {
	var reader services.ConfigReader = services.JSONConfigReader{}
	conf, err := reader.ReadConfig()

	if err != nil {
		log.Fatalln("Failed to read the configuration")
	}
	handler = services.NewHandler(&conf.Hosts)
}

func main() {
	// create a connection to be used by service clients
	conn, err := grpc.Dial("localhost"+constants.ServerPort, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// setup the request handler
	handler.Clients = rpc.InitializeClients(conn)
	handler.Templates = populateTemplates()
	registerHandlers()

	// setup the server and start listening
	server := http.Server{
		Addr: "localhost" + constants.ClientPort,
	}
	log.Println("Client App listening at port", constants.ClientPort)
	log.Fatal(server.ListenAndServe())
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "client/templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}

func registerHandlers() {
	http.Handle("/", http.RedirectHandler("/home", http.StatusPermanentRedirect))
	http.HandleFunc("/home", handler.Routes["home"])
	http.HandleFunc("/containers", handler.Routes["containers"])

	// handlers for static content
	http.Handle("/js/", http.FileServer(http.Dir("client/public")))
	http.Handle("/vendor/", http.FileServer(http.Dir("client/public")))
	http.Handle("/css/", http.FileServer(http.Dir("client/public")))
}
