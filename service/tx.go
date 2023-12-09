package service

import (
	"fmt"

	"github.com/emersonfbarros/backend-challenge-klever/model"
)

type Transaction struct {
	Addresses []Address `json:"addresses"`
	Block     int       `json:"block"`
	TxID      string    `json:"txID"`
}

type Address struct {
	Address string `json:"address"`
	Value   string `json:"value"`
}

func (s *Services) Tx(models model.IModels, txId string) (*Transaction, error, int) {
	txRef, err, httpCode := models.GetTx(fetcher, txId)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error()), httpCode
	}

	extTx := *txRef

	transaction := Transaction{
		Block: extTx.BlockHeight,
		TxID:  txId,
	}

	for _, vin := range extTx.Vin {
		transaction.Addresses = append(transaction.Addresses, Address{
			Address: vin.Addresses[0],
			Value:   vin.Value,
		})
	}

	for _, vout := range extTx.Vout {
		transaction.Addresses = append(transaction.Addresses, Address{
			Address: vout.Addresses[0],
			Value:   vout.Value,
		})
	}

	return &transaction, nil, 0
}
