//+build !test

package main

import (
	"context"
	"golang-api/handler"
	"log"
	"net/http"
)

func main() {
	var m = http.NewServeMux()
	var s = http.Server{Addr: ":8080", Handler: m}

	m.HandleFunc("/hash", handler.HashPassword)
	m.HandleFunc("/shutdown", handler.ExecuteShutdown)
	m.HandleFunc("/stats", handler.ProcessStats)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	if <- handler.IsShuttingDown {
		if err := s.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("Server has shutdown.")
}