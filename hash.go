package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"
)

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
		stats.ProcessRequest(time.Now().Sub(start).Seconds() * 1000)
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Unsupported HTTP method '%v'", r.Method)
	}
}

func hashSHA512(text string) []byte {
	hash := sha512.New()
	hash.Write([]byte(text))
	return hash.Sum(nil)
}