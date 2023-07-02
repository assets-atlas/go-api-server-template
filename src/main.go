package main

import (
	log "github.com/sirupsen/logrus"
	"os/signal"
	"strings"
	"syscall"

	//"log"
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

	logLevel := os.Getenv("LOG_LEVEL")
	switch strings.ToLower(logLevel) {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)

	}

	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})

	log.Printf("server starting")

	hostName, err := os.Hostname()
	if err != nil {
		log.Warn(err)
	}

	log.WithFields(
		log.Fields{
			"server_startup_info": log.Fields{
				"service_name": serviceName,
				"http_port":    httpPort,
				"version":      version,
				"log_level":    logLevel,
				"hostname":     hostName,
			}},
	).Info("server started...")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(`:`+httpPort, router))

	//gracefulShutdown()

}

func gracefulShutdown() {
	// Create a channel to receive the termination signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM)

	// Wait for the termination signal
	<-signalChan

	// Perform any necessary cleanup or shutdown tasks here
	// For example, gracefully stop the server

	log.Println("Server shutting down...")
	// Perform any cleanup or shutdown tasks here

	os.Exit(0)
}
