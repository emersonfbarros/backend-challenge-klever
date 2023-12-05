package model

import (
	"encoding/json"
	"math/big"
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

func Utxo(address string) (*[]UtxoConverted, error) {
	body, err := Fetch("utxo", address)
	if err != nil {
		return nil, err
	}

	var utxoRes []UtxoRes
	if err := json.Unmarshal(body, &utxoRes); err != nil {
		return nil, err
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

	return &utxoConverted, nil
}
