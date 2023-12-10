package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AddressRes struct {
	Balance       string `json:"balance"`
	Txs           int    `json:"txs"`
	TotalSent     string `json:"totalSent"`
	TotalReceived string `json:"totalReceived"`
}

func (handler *Models) Address(fetcher IFetcher, address string) (*AddressRes, error, int) {
	body, err := fetcher.Fetch("address", address)
	if err != nil {
		fmt.Printf("\n\nERRO\n\n")
		return nil, fmt.Errorf("Failed to request external resource"), http.StatusBadGateway
	}
	fmt.Printf("\n\nSEM ERRO\n\n")

	var addressRes AddressRes
	if err := json.Unmarshal(body, &addressRes); err != nil {
		return nil, fmt.Errorf("Internal server error"), http.StatusInternalServerError
	}

	if addressRes.Txs == 0 {
		return nil, fmt.Errorf("Address %s not found", address), http.StatusNotFound
	}

	return &addressRes, nil, 0
}
