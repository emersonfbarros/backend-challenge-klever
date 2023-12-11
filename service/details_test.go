package service

import (
	"errors"
	"net/http"
	"testing"

	"github.com/emersonfbarros/backend-challenge-klever/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mock IServices
type MockServices struct {
	mock.Mock
}

func (m *MockServices) BalanceCalc(models model.IModels, address string) (*BalanceResult, error, int) {
	args := m.Called(models, address)
	return args.Get(0).(*BalanceResult), args.Error(1), args.Int(2)
}

func (m *MockServices) Details(services IServices, models model.IModels, address string) (*AddressInfo, error, int) {
	args := m.Called(services, models, address)
	return args.Get(0).(*AddressInfo), args.Error(1), args.Int(2)
}

func (m *MockServices) Tx(models model.IModels, txId string) (*Transaction, error, int) {
	args := m.Called(models, txId)
	return args.Get(0).(*Transaction), args.Error(1), args.Int(2)
}

func (m *MockServices) Send(models model.IModels, btcTransactionData *SendBtcConverted) (*UtxoNeeded, error, int) {
	args := m.Called(models, btcTransactionData)
	return args.Get(0).(*UtxoNeeded), args.Error(1), args.Int(2)
}

func (m *MockServices) Health(fetcher model.IFetcher) *HealthRes {
	args := m.Called(fetcher)
	return args.Get(0).(*HealthRes)
}

func TestDetailsSuccess(t *testing.T) {
	// creates mocks
	mockServices := new(MockServices)
	mockModels := new(MockIModels)

	// mock expectations
	expectedAddressInfo := &AddressInfo{
		Address:         "mocked_address",
		BalanceExternal: "100",
		TotalTx:         10,
		Balance: BalanceResult{
			Confirmed:   "3492908",
			Unconfirmed: "84699",
		},
		Total: Total{
			Sent:     "50",
			Received: "60",
		},
	}

	// config mocks
	mockServices.On("BalanceCalc", mock.Anything, "mocked_address").
		Return(&expectedAddressInfo.Balance, nil, 0)

	mockModels.On("Address", mock.Anything, "mocked_address").
		Return(&model.AddressRes{
			Balance:       expectedAddressInfo.BalanceExternal,
			Txs:           expectedAddressInfo.TotalTx,
			TotalSent:     expectedAddressInfo.Total.Sent,
			TotalReceived: expectedAddressInfo.Total.Received,
		}, nil, 0)

	// call target method with mocks
	service := &Services{}
	result, err, httpCode := service.Details(mockServices, mockModels, "mocked_address")

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedAddressInfo, result)
	assert.Equal(t, 0, httpCode)

	mockServices.AssertExpectations(t)
	mockModels.AssertExpectations(t)
}

func TestDetailsBalanceCalcError(t *testing.T) {
	mockServices := new(MockServices)
	mockModels := new(MockIModels)

	// mock expectations
	expectedAddressInfo := &AddressInfo{
		Address:         "mocked_address",
		BalanceExternal: "100",
		TotalTx:         10,
		Balance:         BalanceResult{},
		Total: Total{
			Sent:     "50",
			Received: "60",
		},
	}

	// config mocks
	mockServices.On("BalanceCalc", mock.Anything, "mocked_address").
		Return(&expectedAddressInfo.Balance, errors.New("failed to calc balance"), 502)

	mockModels.On("Address", mock.Anything, "mocked_address").
		Return(&model.AddressRes{
			Balance:       expectedAddressInfo.BalanceExternal,
			Txs:           expectedAddressInfo.TotalTx,
			TotalSent:     expectedAddressInfo.Total.Sent,
			TotalReceived: expectedAddressInfo.Total.Received,
		}, nil, 0)

	// call target method with mocks
	service := &Services{}
	result, err, httpCode := service.Details(mockServices, mockModels, "mocked_address")

	// assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "failed to calc balance", err.Error())
	assert.Equal(t, http.StatusBadGateway, httpCode)

	mockModels.AssertExpectations(t)
}

func TestDetailsAddressError(t *testing.T) {
	mockServices := new(MockServices)
	mockModels := new(MockIModels)

	// mock expectations
	expectedAddressInfo := &AddressInfo{
		Address:         "mocked_address",
		BalanceExternal: "100",
		TotalTx:         10,
		Balance:         BalanceResult{},
		Total: Total{
			Sent:     "50",
			Received: "60",
		},
	}

	// config mocks
	mockServices.On("BalanceCalc", mock.Anything, "mocked_address").
		Return(&expectedAddressInfo.Balance, nil, http.StatusBadGateway)

	mockModels.On("Address", mock.Anything, "mocked_address").
		Return(&model.AddressRes{
			Balance:       expectedAddressInfo.BalanceExternal,
			Txs:           expectedAddressInfo.TotalTx,
			TotalSent:     expectedAddressInfo.Total.Sent,
			TotalReceived: expectedAddressInfo.Total.Received,
		}, errors.New("failed to get address"), http.StatusBadGateway)

	// call target method with mocks
	service := &Services{}
	result, err, httpCode := service.Details(mockServices, mockModels, "mocked_address")

	// asserions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "failed to get address", err.Error())
	assert.Equal(t, http.StatusBadGateway, httpCode)

	mockModels.AssertExpectations(t)
}
