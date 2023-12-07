package service

import (
	"math/big"
	"testing"

	"github.com/emersonfbarros/backend-challenge-klever/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSendSuccess(t *testing.T) {
	// creates mocks
	mockModels := new(MockIModels)

	// mock utxos
	mockedUtxoConverted := []model.UtxoConverted{
		{
			Txid:          "jhgavapgbty198g8qy09173urhvcsd",
			Value:         big.NewInt(97324),
			Confirmations: 1,
		},
		{
			Txid:          "rrh0u14979fdbubgfkjlhjo34yu90",
			Value:         big.NewInt(284980),
			Confirmations: 42,
		},
		{
			Txid:          "v9qtyrhb910y53hy9bh0b70w7",
			Value:         big.NewInt(295782),
			Confirmations: 951,
		},
		{
			Txid:          "0dc9abyf7h124037t1828kcjnsdjid",
			Value:         big.NewInt(9134714),
			Confirmations: 3209,
		},
	}

	expectedSendResult := &UtxoNeeded{
		Utxos: []Utxo{
			Utxo{
				Txid:   "0dc9abyf7h124037t1828kcjnsdjid",
				Amount: "9134714",
			},
			Utxo{
				Txid:   "v9qtyrhb910y53hy9bh0b70w7",
				Amount: "295782",
			},
		},
	}

	// config mocks
	mockModels.On("Utxo", mock.Anything, "mocked_address").
		Return(&mockedUtxoConverted, nil)

	// mock param
	btcTransaction := &SendBtcConverted{
		Address: "mocked_address",
		Amount:  big.NewInt(9182416),
	}

	service := &Services{}
	result, err := service.Send(mockModels, btcTransaction)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedSendResult, result)

	mockModels.AssertExpectations(t)
}
