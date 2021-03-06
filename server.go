package hummingbird

import (
	"log"
	"net/http"
	"strconv"
)

// RegisterAPIEndpoints initializes API Endpoints
func RegisterAPIEndpoints() {
	http.HandleFunc("/", renderFullyRequestDetails)
	http.HandleFunc("/method", renderRequestMethod)
	http.HandleFunc("/headers", renderHeaders)
	http.HandleFunc("/body", renderBody)
	http.HandleFunc("/remote_addr", renderRemoteAddr)
}

// Run starts up a HTTP server
func Run() {
	_, err := strconv.Atoi(ServerConfig.Port)
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(":"+ServerConfig.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
