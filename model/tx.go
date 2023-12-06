package model

import "encoding/json"

type vin struct {
	Addresses []string `json:"addresses"`
	Value     string   `json:"value"`
}

type vout struct {
	Value     string   `json:"value"`
	Addresses []string `json:"addresses"`
}

type ExtTx struct {
	Vin         []vin  `json:"vin"`
	Vout        []vout `json:"vout"`
	BlockHeight int    `json:"blockHeight"`
}

func (handler *Models) GetTx(fetcher IFetcher, txId string) (*ExtTx, error) {
	body, err := fetcher.Fetch("tx", txId)
	if err != nil {
		return nil, err
	}

	var extTx ExtTx
	if err := json.Unmarshal(body, &extTx); err != nil {
		return nil, err
	}

	return &extTx, nil
}
