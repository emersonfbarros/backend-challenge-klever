package model

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
)

type UtxoRes struct {
	Txid          string `json:"txid"`
	Value         string `json:"value"`
	Confirmations int    `json:"confirmations"`
}

type UtxoConverted struct {
	Txid          string
	Value         *big.Int
	Confirmations int
}

func (handler *Models) Utxo(fetcher IFetcher, address string) (*[]UtxoConverted, error, int) {
	body, err := fetcher.Fetch("utxo", address)
	if err != nil {
		return nil, fmt.Errorf("Failed to request external resource"), http.StatusBadGateway
	}

	var bodyVerification struct {
		Error string `json:"error"`
	}

	json.Unmarshal(body, &bodyVerification)
	if len(bodyVerification.Error) != 0 {
		return nil, fmt.Errorf("Address %s not found", address), http.StatusNotFound
	}

	var utxoRes []UtxoRes
	if err := json.Unmarshal(body, &utxoRes); err != nil {
		return nil, fmt.Errorf("Internal server error"), http.StatusInternalServerError
	}

	var utxoConverted []UtxoConverted
	for _, utxo := range utxoRes {
		valueBigInt := new(big.Int)
		valueBigInt.SetString(utxo.Value, 10)

		utxoConverted = append(utxoConverted, UtxoConverted{
			Txid:          utxo.Txid,
			Value:         valueBigInt,
			Confirmations: utxo.Confirmations,
		})
	}

	return &utxoConverted, nil, 0
}
