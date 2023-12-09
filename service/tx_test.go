package service

import (
	"errors"
	"net/http"
	"testing"

	"github.com/emersonfbarros/backend-challenge-klever/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mock IModel
type MockIModels struct {
	mock.Mock
}

func (m *MockIModels) GetTx(fetcher model.IFetcher, txId string) (*model.ExtTx, error, int) {
	args := m.Called(fetcher, txId)
	return args.Get(0).(*model.ExtTx), args.Error(1), args.Int(2)
}

func (m *MockIModels) Utxo(fetcher model.IFetcher, address string) (*[]model.UtxoConverted, error, int) {
	args := m.Called(fetcher, address)
	return args.Get(0).(*[]model.UtxoConverted), args.Error(1), args.Int(2)
}

func (m *MockIModels) Address(fetcher model.IFetcher, address string) (*model.AddressRes, error, int) {
	args := m.Called(fetcher, address)
	return args.Get(0).(*model.AddressRes), args.Error(1), args.Int(2)
}

func TestTxSuccess(t *testing.T) {
	modelsMock := new(MockIModels)

	// expected data
	expectedTxId := "someTxId"
	expectedExtTx := &model.ExtTx{
		BlockHeight: 675674,
		Vin: []model.Vin{
			{Addresses: []string{"bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r"}, Value: "484817655"},
		},
		Vout: []model.Vout{
			{Addresses: []string{"36iYTpBFVZPbcyUs8pj3BtutZXzN6HPNA6"}, Value: "623579"},
			{Addresses: []string{"bc1qe29ydjtwyjdmffxg4qwtd5wfwzdxvnap989glq"}, Value: "3283266"},
			{Addresses: []string{"bc1qanhueax8r4cn52r38f2h727mmgg6hm3xjlwd0x"}, Value: "90311"},
		},
	}

	modelsMock.On("GetTx", mock.Anything, expectedTxId).Return(expectedExtTx, nil, 0)

	services := Services{}

	result, err, httpCode := services.Tx(modelsMock, expectedTxId)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedExtTx.BlockHeight, result.Block)
	assert.Equal(t, expectedTxId, result.TxID)
	assert.Equal(t, "bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r", result.Addresses[0].Address)
	assert.Equal(t, "484817655", result.Addresses[0].Value)
	assert.Equal(t, "36iYTpBFVZPbcyUs8pj3BtutZXzN6HPNA6", result.Addresses[1].Address)
	assert.Equal(t, "623579", result.Addresses[1].Value)
	assert.Equal(t, "bc1qe29ydjtwyjdmffxg4qwtd5wfwzdxvnap989glq", result.Addresses[2].Address)
	assert.Equal(t, "3283266", result.Addresses[2].Value)
	assert.Equal(t, "bc1qanhueax8r4cn52r38f2h727mmgg6hm3xjlwd0x", result.Addresses[3].Address)
	assert.Equal(t, "90311", result.Addresses[3].Value)
	assert.Equal(t, 0, httpCode)

	modelsMock.AssertExpectations(t)
}

func TestTxError(t *testing.T) {
	modelsMock := new(MockIModels)

	expectedExtTx := &model.ExtTx{
		BlockHeight: 123,
		Vin:         []model.Vin{{Addresses: []string{"address1"}, Value: "value1"}},
		Vout:        []model.Vout{{Addresses: []string{"address2"}, Value: "value2"}},
	}

	expectedTxId := "someTxId"

	modelsMock.On("GetTx", mock.Anything, expectedTxId).
		Return(expectedExtTx, errors.New("failed to request external resource"), http.StatusBadGateway)

	services := Services{}

	result, err, httpCode := services.Tx(modelsMock, expectedTxId)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "failed to request external resource", err.Error())
	assert.Equal(t, http.StatusBadGateway, httpCode)

	modelsMock.AssertExpectations(t)
}
