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

func (s *Services) Tx(models model.IModels, txId string) (*Transaction, error) {
	txRef, err := models.GetTx(fetcher, txId)
	if err != nil {
		logger.Errorf("failed to unmarshal api response %v", err.Error())
		return nil, fmt.Errorf("failed to request external resouce")
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

	return &transaction, nil
}
