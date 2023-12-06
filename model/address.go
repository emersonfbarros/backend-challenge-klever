package model

import "encoding/json"

type AddressRes struct {
	Balance       string `json:"balance"`
	Txs           int    `json:"txs"`
	TotalSent     string `json:"totalSent"`
	TotalReceived string `json:"totalReceived"`
}

func (handler *Models) Address(fetcher Fetcher, address string) (*AddressRes, error) {
	body, err := fetcher.Fetch("address", address)
	if err != nil {
		return nil, err
	}

	var addressRes AddressRes
	if err := json.Unmarshal(body, &addressRes); err != nil {
		return nil, err
	}

	return &addressRes, nil
}
