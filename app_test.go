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
func TestHealthCheck(t *testing.T) {
	startUp()
	req, _ := http.NewRequest("GET", "/api/v1/health-check", nil)
	response := httptest.NewRecorder()
	testRouter := mux.NewRouter()
	testRouter.HandleFunc("/api/v1/health-check", healthCheck.GetHealth).Methods("GET")
	testRouter.ServeHTTP(response, req)
	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "{\"Message\":\"Ah Ah Ah Ah Stayin Alive!!\",\"Data\":null}" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}
