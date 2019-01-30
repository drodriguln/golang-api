package handler

import (
	"crypto/sha256"
	"crypto/sha512"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestHashPassword(t *testing.T) {
	req, err := http.NewRequest("POST", "/hash", strings.NewReader("password=angryMonkey"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/form-data")
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(HashPassword)
	handler.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Handler unexpectedly returned status code %v instead of %v", status, http.StatusOK)
	}

	expected := "z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg/SpIdNs6c5H0NE8XYXysP+DGNKHfuwvY7kxvUdBeoGlODJ6+SfaPg=="
	if r.Body.String() != expected {
		t.Errorf("Handler unexpectedly returned body %v instead of %v", r.Body.String(), expected)
	}
}

func TestHashPasswordUnsupportedMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/hash", nil)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(HashPassword)
	handler.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusBadRequest {
		t.Errorf("Handler unexpectedly returned status code %v instead of %v", status, http.StatusBadRequest)
	}
}

func TestHashSHA512(t *testing.T) {
	input := "angryMonkey"
	hash := sha512.New()
	hash.Write([]byte(input))
	expected := hash.Sum(nil)
	actual := hashSHA512(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("The expected and actual SHA512 hashed bytes for the input string '%v' do not match.", input)
	}
}

func TestHashSHA512Fail(t *testing.T) {
	input := "angryMonkey"
	hash := sha256.New()
	hash.Write([]byte(input))
	expected := hash.Sum(nil)
	actual := hashSHA512(input)
	if reflect.DeepEqual(actual, expected) {
		t.Errorf("The expected and actual SHA512 hashed bytes for the input string '%v' were unexpectedly equal.", input)
	}
}
