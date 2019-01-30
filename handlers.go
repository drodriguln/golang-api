package main

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var stats = HashRequestStats{}
var mutex = &sync.Mutex{}
var shutdownCh = make(chan bool)

// statsHandler takes a given HTTP request and returns the statistics information relevent to the number of previous requests made to hash a password.
func statsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		data, err := json.Marshal(stats)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(data); err != nil {
			log.Fatal(err)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Unsupported HTTP method '%v'", r.Method)
	}
}

// hashHandler takes a given HTTP request, and encodes a password string from form data attached to the request into base64, and then hashes in SHA512. Request statistics are also stored in a global struct.
// It returns the base64 encoded and SHA512 hashed password after a 5 seconds pause.
func hashHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	time.Sleep(time.Duration(5) * time.Second)
	switch r.Method {
	case http.MethodPost:
		password := r.FormValue("password")
		passwordHashed := hashSHA512(password)
		passwordHashedAndEncoded := base64.StdEncoding.EncodeToString(passwordHashed)
		if _, err := fmt.Fprintf(w, passwordHashedAndEncoded); err != nil {
			log.Fatal(err)
		}
		elapsedTime := time.Now().Sub(start).Seconds() * 1000
		stats.Process(int(elapsedTime))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Unsupported HTTP method '%v'", r.Method)
	}
}

// Takes a given HTTP request, and gracefully shuts down the server.
// It uses a channel to keep the server alive until any remaining active connections are completed before shutting down.
func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if _, err := fmt.Fprintf(w, "Server has shutdown."); err != nil {
			log.Fatal(err)
		}
		shutdownCh <- true
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Unsupported HTTP method '%v'", r.Method)
	}
}

// hashSHA512 returns a text string in a SHA512 hashed form.
func hashSHA512(text string) []byte {
	hash := sha512.New()
	hash.Write([]byte(text))
	return hash.Sum(nil)
}