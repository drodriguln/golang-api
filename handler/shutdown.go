package handler

import (
	"fmt"
	"log"
	"net/http"
)

var IsShuttingDown = make(chan bool)

// ExecuteShutdown takes a given HTTP request, and gracefully shuts down the server.
// It uses a channel to keep the server alive until any remaining active connections are completed before shutting down.
func ExecuteShutdown(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if _, err := fmt.Fprintf(w, "Server is gracefully shutting down."); err != nil {
			log.Fatal(err)
		}
		IsShuttingDown <- true

	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Unsupported HTTP method '%v'", r.Method)
	}
}