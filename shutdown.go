package main

import (
	"fmt"
	"log"
	"net/http"
)

var shutdownCh = make(chan bool)

func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Server has shutdown."); err != nil {
		log.Fatal(err)
	}
	shutdownCh <- true
}