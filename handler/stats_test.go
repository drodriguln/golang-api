package handler

import (
	"encoding/json"
	"golang-api/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProcessStats(t *testing.T) {
	req, err := http.NewRequest("GET", "/stats", nil)
	if err != nil {
		t.Fatal(err)
	}

	model.Stats = model.HashRequestStats{1, 5001,5001}
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(ProcessStats)
	handler.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Handler unexpectedly returned status code %v instead of %v", status, http.StatusOK)
	}

	expected, err := json.Marshal(model.Stats)
	if r.Body.String() != string(expected) {
		t.Errorf("Handler unexpectedly returned body %v instead of %v", r.Body.String(), string(expected))
	}
}

func TestProcessStatsUnsupportedMethod(t *testing.T) {
	req, err := http.NewRequest("HEAD", "/stats", nil)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(ProcessStats)
	handler.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusBadRequest {
		t.Errorf("Handler unexpectedly returned status code %v instead of %v", status, http.StatusBadRequest)
	}
}