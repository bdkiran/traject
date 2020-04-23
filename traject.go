package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/bdkiran/traject/api"
	"github.com/bdkiran/traject/persist"
)

type config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		URI      string `yaml:"uri"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	}
}

func main() {
	//Instaciats a new config object, getting data from config file
	var cfg config
	getConfigurations(&cfg)

	serveURL := cfg.Server.Host + ":" + cfg.Server.Port
	log.Println("Staring Traject at:")
	log.Println("http://" + serveURL)

	persist.ConnectToMongo(cfg.Database.URI, cfg.Database.Username, cfg.Database.Password)

	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(serveURL, router))
}

func getConfigurations(cfg *config) {
	const configFile = "config.yaml"
	f, err := os.Open(configFile)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
