package main

import (
	"net/http"
	"time"

	"github.com/bdkiran/traject/api"
	"github.com/bdkiran/traject/persist"
	"github.com/bdkiran/traject/utils"
)

func main() {
	//Initializes a new logger to be used throughout the application
	//Currently hard coded logger configurations
	utils.LoggerInit()

	//Obtains configurations for the applications
	var cfg utils.Config
	utils.GetConfigurations(&cfg)

	serveURL := "localhost:" + cfg.Server.Port
	port := ":" + cfg.Server.Port
	utils.DefaultLogger.Info.Println("You can now view form-app in the browser.")
	utils.DefaultLogger.Info.Println("Local: http://" + serveURL)

	//Initilizes a new kafka producer
	persist.InitilizeProducer(cfg.Kafka.BootstrapServer)

	//Create a new connection to mongo.
	// err := persist.InitializeMongo(cfg.Database.URI, cfg.Database.Username, cfg.Database.Password)
	// if err != nil {
	// 	utils.ProcessError(err)
	// }

	//Initializes all routers and creats a listener
	router := api.NewRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	utils.DefaultLogger.Error.Fatal(srv.ListenAndServe())
}
