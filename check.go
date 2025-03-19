package main

import (
	"fmt"
	"net/http"
)

// Check performs an HTTP GET request to the specified domain and port
func Check(domain, port string) string {
	url := fmt.Sprintf("http://%s:%s", domain, port)

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Sprintf("Failed to create request: %v", err)
	}

	// Set a common user-agent (mimicking a browser)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("Failed to reach %s:%s - %v", domain, port, err)
	}
	defer resp.Body.Close()

	return fmt.Sprintf("Status: %s (%d)", resp.Status, resp.StatusCode)
}
