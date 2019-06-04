package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gauravgahlot/dockerdoodle/client/controller"
	"github.com/gauravgahlot/dockerdoodle/client/helpers"
	"github.com/gauravgahlot/dockerdoodle/client/rpc"
	"github.com/gauravgahlot/dockerdoodle/constants"
	"github.com/gauravgahlot/dockerdoodle/types"

	"google.golang.org/grpc"
)

var useLocal = flag.Bool("L", false, "use localhost as the only Docker Host")

func main() {
	flag.Parse()

	conn, err := grpc.Dial("localhost"+constants.ServerPort, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err.Error())
	}

	clients := rpc.InitializeClients(conn)
	templates := populateTemplates()

	var config *types.Config
	if *useLocal {
		config = configForLocalEnv()
	} else {
		config = readConfiguration()
	}

	controller.Startup(templates, clients, &config.Hosts)
	server := http.Server{
		Addr: constants.ClientPort,
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

func readConfiguration() *types.Config {
	var reader helpers.ConfigReader = helpers.JSONConfigReader{}
	conf, err := reader.ReadConfig()
	if err != nil {
		log.Panic("Failed to read the configuration")
	}
	return &conf
}

func configForLocalEnv() *types.Config {
	host, _ := os.Hostname()
	const localIP = "0.0.0.0"
	return &types.Config{
		Hosts: []types.Host{
			types.Host{Name: host, IP: localIP},
		},
	}
}
