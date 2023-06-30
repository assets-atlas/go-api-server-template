package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	version := os.Getenv("SERVER_VERSION")
	if version == "" {
		version = "0.1"
	}

	serviceName := os.Getenv("SERVER_SERVICE_NAME")
	if serviceName == "" {
		serviceName = "example-service"
	}

	httpPort := os.Getenv("SERVER_HTTP_PORT")
	if httpPort == "" {
		httpPort = "2000"
	}

	log.Printf("starting %s\nversion: %s\nhttp port: %s", serviceName, version, httpPort)

	router := NewRouter()

	log.Fatal(http.ListenAndServe(`:`+httpPort, router))
}
