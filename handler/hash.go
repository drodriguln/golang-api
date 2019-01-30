package handler

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"golang-api/model"
	"log"
	"net/http"
	"time"
)

// HashPassword takes a given HTTP request, and encodes a password string from form data attached to the request into base64, and then hashes in SHA512. Request statistics are also stored in a global struct.
// It returns the base64 encoded and SHA512 hashed password after a 5 seconds pause.
func HashPassword(w http.ResponseWriter, r *http.Request) {
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
		model.Stats.Process(int(elapsedTime))
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
