package main

import (
	"context"
	"log"
	"net/http"
)

var m = http.NewServeMux()
var s = http.Server{Addr: ":8080", Handler: m}

func main() {
	m.HandleFunc("/hash", hashHandler)
	m.HandleFunc("/shutdown", shutdownHandler)
	m.HandleFunc("/stats", statsHandler)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	if (<-shutdownCh){
		if err := s.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("Server has shutdown.")
}