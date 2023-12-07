package model

import (
	"fmt"
	"io"
	"net/http"
)

// route must be "address", or "utxo", ou "tx"
func (client *Fetcher) Fetch(route string, value string) ([]byte, error) {
	if route != "address" && route != "utxo" && route != "tx" {
		return nil, fmt.Errorf("Invalid route parameter: %s", route)
	}

	apiUrl := fmt.Sprintf("%s/%s/%s", client.BaseURL, route, value)

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		logger.Errorf("failed to create http request %v", err.Error())
		return nil, err
	}

	req.SetBasicAuth(client.Username, client.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Errorf("failed to make http request %v", err.Error())
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
