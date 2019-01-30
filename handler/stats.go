package handler

import (
	"encoding/json"
	"golang-api/model"
	"log"
	"net/http"
)

// ProcessStats takes a given HTTP request and returns the statistics information relevent to the number of previous requests made to hash a password.
func ProcessStats(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		data, err := json.Marshal(model.Stats)
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
