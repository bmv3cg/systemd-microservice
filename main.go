package main

import (
		"log"
		"net/http"
	   gmux "github.com/bmv3cg/systemd-microservice/cmd/api"
	)

func main() {

	router := gmux.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
