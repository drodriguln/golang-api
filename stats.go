package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

var stats = RequestStats{}
var mutex = &sync.Mutex{}

type RequestStats struct {
	HashRequests     int     `json:"total"`
	AverageTime      int     `json:"average"`
	totalTime        float64
}

func (s *RequestStats) ProcessRequest(time float64) {
	mutex.Lock()
	s.HashRequests++
	s.totalTime += time
	s.AverageTime = int(s.totalTime) / s.HashRequests
	mutex.Unlock()
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
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
}