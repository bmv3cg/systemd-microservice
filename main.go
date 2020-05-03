package main

import (
	gmux "github.com/bmv3cg/systemd-microservice/cmd/api"
	"github.com/bmv3cg/systemd-microservice/pkg/config"
	"log"
	"net/http"
)

func main() {

	// Initalise configuration file
	config.InitViper()

	// configure gorilla mux router
	router := gmux.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
