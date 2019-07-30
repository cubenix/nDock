package helpers

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/gauravgahlot/dockerdoodle/pkg/constants"
	"github.com/gauravgahlot/dockerdoodle/pkg/types"
)

// PopulateTemplates populatest the available templates
func PopulateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "app/templates"
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

// ReadConfiguration reads the configuration from JSON config file
func ReadConfiguration() *types.Config {
	conf, err := readJSONConfig()
	if err != nil {
		log.Panic("Failed to read the configuration")
	}
	return &conf
}

// ConfigForLocalEnv sets the localhost as the only Docker Host
func ConfigForLocalEnv() *types.Config {
	host, _ := os.Hostname()
	return &types.Config{
		Hosts: []types.Host{
			types.Host{Name: host, IP: constants.LocalIP},
		},
	}
}

func readJSONConfig() (types.Config, error) {
	var config types.Config

	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln(err)
		return config, err
	}

	jsonErr := json.Unmarshal(data, &config)
	if jsonErr != nil {
		log.Fatalln("Invalid JSON file:", err)
		return config, err
	}
	return config, nil
}
