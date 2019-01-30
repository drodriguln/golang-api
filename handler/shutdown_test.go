package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExecuteShutdownUnsupportedMethod(t *testing.T) {
	req, err := http.NewRequest("PATCH", "/shutdown", nil)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(ExecuteShutdown)
	handler.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusBadRequest {
		t.Errorf("Handler unexpectedly returned status code %v instead of %v", status, http.StatusBadRequest)
	}
}
