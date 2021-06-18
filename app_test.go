package main

import (
	"net/http"
	"net/http/httptest"
	"secretSanta/api/healthCheck"
	"testing"

	"github.com/gorilla/mux"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func checkResponseBody(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected response code %s. Got %s\n", expected, actual)
	}
}

func TestHealthCheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/health-check", nil)
	response := httptest.NewRecorder()
	testRouter := mux.NewRouter()
	testRouter.HandleFunc("/api/v1/health-check", healthCheck.GetHealth).Methods("GET")
	testRouter.ServeHTTP(response, req)
	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponseBody(t, "{\"Message\":\"Ah Ah Ah Ah Stayin Alive!!\",\"Data\":null}", response.Body.String())
}
