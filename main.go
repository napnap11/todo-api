package main

import (
	"flag"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/napnap11/todo-api/internal/app/routes"
	"github.com/napnap11/todo-api/internal/pkg/configs"
	"github.com/napnap11/todo-api/internal/pkg/util/http"
)

func main() {

	var port *string
	var configPath *string

	for _, envVar := range os.Environ() {
		variable := strings.Split(envVar, "=")
		if variable[0] == "PORT" {
			port = &variable[1]
		}
	}
	if port == nil {
		port = flag.String("p", "8080", "port number")
	}

	configPath = flag.String("config", "configs", "configs path")
	if err := configs.InitAppConfig(*configPath); err != nil {
		log.Panic(err)
	}

	router := routes.NewRouter()
	server := http.NewServer(*port, router)

	log.Infof("==== Server had been started ====")
	server.ListenAndServeWithGracefulShutdown(configs.AppConfig.Server.GracefulShutdownTime)
	log.Infof("==== Server had been stopped ====")
}
