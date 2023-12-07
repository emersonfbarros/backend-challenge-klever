package service

import (
	"errors"
	"math/big"
	"testing"

	"github.com/emersonfbarros/backend-challenge-klever/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBalanceCalcSuccess(t *testing.T) {
	// creates mocks
	mockModels := new(MockIModels)

	// mock expectations
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

	expectedBalanceCalcResult := &BalanceResult{
		Confirmed:   "9715476",
		Unconfirmed: "97324",
	}

	// config mocks
	mockModels.On("Utxo", mock.Anything, "mocked_uxto").
		Return(&mockedUtxoConverted, nil)

	service := &Services{}
	result, err := service.BalanceCalc(mockModels, "mocked_uxto")

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedBalanceCalcResult, result)

	mockModels.AssertExpectations(t)
}

func TestBalanceCalcError(t *testing.T) {
	// creates mocks
	InitService()
	originalLogger := logger // keeps orginal logger saved
	mockLogger := new(MockLogger)
	logger = mockLogger // replaces original with mock
	defer func() {
		logger = originalLogger // restores the original logger after test
	}()
	mockModels := new(MockIModels)

	// mock expectations
	mockedUtxoConverted := []model.UtxoConverted{
		{
			Txid:          "whatever",
			Value:         big.NewInt(100),
			Confirmations: 1,
		},
	}

	// config mocks
	mockModels.On("Utxo", mock.Anything, "mocked_uxto").
		Return(&mockedUtxoConverted, errors.New("failed to get utxos"))

	mockLogger.On("Errorf", mock.Anything, mock.Anything).Return()

	service := &Services{}
	result, err := service.BalanceCalc(mockModels, "mocked_uxto")

	// asserions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "failed to request external resource", err.Error())

	mockLogger.AssertCalled(t, "Errorf", "failed to unmarshal api response %v", mock.Anything)
	mockModels.AssertExpectations(t)
}
