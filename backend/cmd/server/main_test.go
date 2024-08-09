package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/ovrdos/goforex/backend/pkg/handlers"
)

func TestLLMHandler(t *testing.T) {
	// Set up the HTTP server with the handler
	mux := http.NewServeMux()
	mux.HandleFunc("/llm", handlers.LLMHandler)

	// Create a test server
	ts := httptest.NewServer(mux)
	defer ts.Close()

	// Create a request body for the test
	requestBody := strings.NewReader(`{"input":"Hello"}`)

	// Make a POST request to the /llm endpoint
	resp, err := http.Post(ts.URL+"/llm", "application/json", requestBody)
	if err != nil {
		t.Fatalf("Failed to make POST request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	// Additional checks can be made here, such as checking the response body
	// For example, decoding the JSON response and checking the "output" field
}

