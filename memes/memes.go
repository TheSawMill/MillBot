package memes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const MemseBaseUrl = "https://api.giphy.com/v1/gifs/random"

func GetRandomMeme(apiKey string) (string, error) {
	apiParam := "api_key"

	uri := MemseBaseUrl + "?" + apiParam + "=" + apiKey

	resp, err := http.Get(uri)

	if err != nil {
		log.Println("Error conducting random meme get request:", err)
	}

	// Ensure the response boyd is closed to prevent resource leaks
	defer resp.Body.Close()

	// Check the respons status code
	if resp.StatusCode != http.StatusOK {
		log.Printf("Request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error reading response body:", err)
	}

	// Unmarshal into generic map
	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Error unmarshaling out gify response", err)
	}

	data, ok := result["data"].(map[string]any)

	if !ok {
		log.Println("Data field not found in json response")
	}

	url, ok := data["url"].(string)

	if !ok {
		log.Println("Url field not found in json response")
	}

	return url, nil
}
