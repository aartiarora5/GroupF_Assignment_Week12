package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
)

func TestGetPrice(t *testing.T) {
	// Create a request to the "/price" endpoint
	req, err := http.NewRequest("GET", "/price", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create an HTTP handler using the getPrice function
	handler := http.HandlerFunc(getPrice)

	// Serve the HTTP request to the ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Parse JSON response
	var response map[string]float64
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	// Check if the required keys are present in the response
	keys := []string{"bitcoin", "ethereum", "tether"}
	for _, key := range keys {
		if _, ok := response[key]; !ok {
			t.Errorf("Response does not contain key: %v", key)
		}
	}
}