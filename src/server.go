package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type DefaultRouteResponse struct {
	ServiceName string `json:"service_name"`
	Version     string `json:"version"`
}

func DefaultRoute() http.HandlerFunc {

	version := os.Getenv("SERVER_VERSION")
	if version == "" {
		version = "0.1"
	}

	serviceName := os.Getenv("SERVER_SERVICE_NAME")
	if serviceName == "" {
		serviceName = "example-service"
	}

	return func(w http.ResponseWriter, r *http.Request) {
		resp := DefaultRouteResponse{
			ServiceName: serviceName,
			Version:     version,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}

}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", DefaultRoute()).Methods("GET")

	return r
}
