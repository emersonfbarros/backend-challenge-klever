package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Vin struct {
	Addresses []string `json:"addresses"`
	Value     string   `json:"value"`
}

type Vout struct {
	Value     string   `json:"value"`
	Addresses []string `json:"addresses"`
}

type ExtTx struct {
	Vin         []Vin  `json:"vin"`
	Vout        []Vout `json:"vout"`
	BlockHeight int    `json:"blockHeight"`
}

func (handler *Models) GetTx(fetcher IFetcher, txId string) (*ExtTx, error, int) {
	body, err := fetcher.Fetch("tx", txId)
	if err != nil {
		return nil, err, http.StatusBadGateway
	}

	var extTx ExtTx
	if err := json.Unmarshal(body, &extTx); err != nil {
		return nil, err, http.StatusInternalServerError
	}

	if extTx.BlockHeight == 0 {
		return nil, fmt.Errorf("Transaction %s not found", txId), http.StatusNotFound
	}

	return &extTx, nil, 0
}
