package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelathCheck(t *testing.T){
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheck)
	handler.ServeHTTP(w, req)
	status := w.Code
	if status != http.StatusOK {
		t.Errorf("handler return wrong status code : got %v want %v", status, http.StatusOK)
	}

	expected := `{"Status":"Good"}`
	if w.Body.String() != expected {
		t.Errorf("handler response unxpected body: got %v want %v", w.Body.String(), expected)
	}
}