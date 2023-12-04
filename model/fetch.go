package model

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// route must be "address", or "utxo", ou "tx"
func Fetch(route string, value string) ([]byte, error) {
	if route != "address" && route != "utxo" && route != "tx" {
		return nil, fmt.Errorf("Invalid route parameter: %s", route)
	}

	baseUrl := os.Getenv("BASE_URL")
	apiUrl := fmt.Sprintf("%s/%s/%s", baseUrl, route, value)

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		logger.Errorf("failed to create http request %v", err.Error())
		return nil, err
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("failed to make http reques %v", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("Failed to read request body %v", err.Error())
		return nil, err
	}

	return body, nil
}
