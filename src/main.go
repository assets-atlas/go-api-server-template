package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os/signal"
	"strings"
	"syscall"
	"time"

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

	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		log.Fatal("DATABASE_NAME environment variable must be set")
	}

	dbUser := os.Getenv("DATABASE_USER")
	if dbName == "" {
		log.Fatal("DATABASE_USER environment variable must be set")
	}

	dbPassword := os.Getenv("DATABASE_PASSWORD")
	if dbName == "" {
		log.Fatal("DATABASE_PASSWORD environment variable must be set")
	}

	dbHost := os.Getenv("DATABASE_HOST")
	if dbName == "" {
		log.Fatal("DATABASE_HOST environment variable must be set")
	}

	dbPort := os.Getenv("DATABASE_PORT")
	if dbName == "" {
		log.Fatal("DATABASE_PORT environment variable must be set")
	}

	dbSslMode := os.Getenv("DATABASE_SSL_MODE")
	if dbName == "" {
		log.Info("DATABASE_SSL_MODE not set. defaulting to require")
		dbSslMode = "require"
	}

	connStr := "user=" + dbUser + " password=" + dbPassword + " host=" + dbHost + " port=" + dbPort + " dbname=" + dbName + " sslmode=" + dbSslMode + " sslrootcert=/Users/rbarnes/Downloads/ca-certificate.crt"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(10 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Errorf("Error establishing database connection: %s", err)
		return
	}

	log.Println("Database connection established")

	vaultAddr := os.Getenv("VAULT_ADDR")
	if vaultAddr == "" {
		log.Fatal("VAULT_ADDR has not been set")
	}

	vaultToken := os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		log.Fatal("VAULT_TOKEN has not been set")
	}

	keyName := os.Getenv("VAULT_TRANSIT_KEY_NAME")
	if keyName == "" {
		log.Fatal("VAULT_TRANSIT_KEY_NAME not set")
	}

	vc := vlt.NewClient(vaultAddr, string(vaultToken))
	vcw := vaultidentity.VaultClientWrapper{VC: &vc}

	router := NewRouter(db, &vc, vcw)

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
