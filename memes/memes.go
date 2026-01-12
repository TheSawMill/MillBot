package memes

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func TestGetRequest() {
	// Make a GET request to an example API endpoint
	resp, err := http.Get("https://api.restful-api.dev/objects")

	if err != nil {
		log.Fatal("Error making HTTP request:", err)
	}

	// Ensure the response body is closed to prevent resource leaks
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	// print the response body as a string
	fmt.Println(string(body))
}
