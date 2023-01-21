package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/nacl/secretbox"
)

func main() {
	// Set the API endpoint URL
	url := "https://chat.openai.com/chat/72567136-802e-497f-accd-da74a0681709"

	// Set the API key
	apiKey := "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"

	// Create a new http client
	client := &http.Client{}

	// Create a new http request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add the API key to the request headers
	req.Header.Add("Authorization", "Bearer "+apiKey)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Close the response body
	defer resp.Body.Close()

	// Print the response body
	fmt.Println(string(body))
}