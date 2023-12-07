package service

import (
	"errors"
	"testing"

	"github.com/emersonfbarros/backend-challenge-klever/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mock IServices
type MockServices struct {
	mock.Mock
}

func (m *MockServices) BalanceCalc(models model.IModels, address string) (*BalanceResult, error) {
	args := m.Called(models, address)
	return args.Get(0).(*BalanceResult), args.Error(1)
}

func (m *MockServices) Details(services IServices, models model.IModels, address string) (*AddressInfo, error) {
	args := m.Called(services, models, address)
	return args.Get(0).(*AddressInfo), args.Error(1)
}

func (m *MockServices) Tx(models model.IModels, txId string) (*Transaction, error) {
	args := m.Called(models, txId)
	return args.Get(0).(*Transaction), args.Error(1)
}

func (m *MockServices) Send(models model.IModels, btcTransactionData *SendBtcConverted) (*UtxoNeeded, error) {
	args := m.Called(models, btcTransactionData)
	return args.Get(0).(*UtxoNeeded), args.Error(1)
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
		Return(&expectedAddressInfo.Balance, nil)

	mockModels.On("Address", mock.Anything, "mocked_address").
		Return(&model.AddressRes{
			Balance:       expectedAddressInfo.BalanceExternal,
			Txs:           expectedAddressInfo.TotalTx,
			TotalSent:     expectedAddressInfo.Total.Sent,
			TotalReceived: expectedAddressInfo.Total.Received,
		}, nil)

	// call target method with mocks
	service := &Services{}
	result, err := service.Details(mockServices, mockModels, "mocked_address")

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedAddressInfo, result)

	mockServices.AssertExpectations(t)
	mockModels.AssertExpectations(t)
}

func TestDetailsBalanceCalcError(t *testing.T) {
	// creates mocks
	InitService()
	originalLogger := logger // keeps orginal logger saved
	mockLogger := new(MockLogger)
	logger = mockLogger // replaces original with mock
	defer func() {
		logger = originalLogger // restores the original logger after test
	}()
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
		Return(&expectedAddressInfo.Balance, errors.New("failed to calc balance"))

	mockModels.On("Address", mock.Anything, "mocked_address").
		Return(&model.AddressRes{
			Balance:       expectedAddressInfo.BalanceExternal,
			Txs:           expectedAddressInfo.TotalTx,
			TotalSent:     expectedAddressInfo.Total.Sent,
			TotalReceived: expectedAddressInfo.Total.Received,
		}, nil)

	mockLogger.On("Errorf", mock.Anything, mock.Anything).Return()

	// call target method with mocks
	service := &Services{}
	result, err := service.Details(mockServices, mockModels, "mocked_address")

	// asserions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "failed to request external resource", err.Error())

	mockLogger.AssertCalled(t, "Errorf", "failed to request address or utxo", mock.Anything)
	mockModels.AssertExpectations(t)
}

func TestDetailsAddressError(t *testing.T) {
	// creates mocks
	InitService()
	originalLogger := logger // keeps orginal logger saved
	mockLogger := new(MockLogger)
	logger = mockLogger // replaces original with mock
	defer func() {
		logger = originalLogger // restores the original logger after test
	}()
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
		Return(&expectedAddressInfo.Balance, nil)

	mockModels.On("Address", mock.Anything, "mocked_address").
		Return(&model.AddressRes{
			Balance:       expectedAddressInfo.BalanceExternal,
			Txs:           expectedAddressInfo.TotalTx,
			TotalSent:     expectedAddressInfo.Total.Sent,
			TotalReceived: expectedAddressInfo.Total.Received,
		}, errors.New("failed to get address"))

	mockLogger.On("Errorf", mock.Anything, mock.Anything).Return()

	// call target method with mocks
	service := &Services{}
	result, err := service.Details(mockServices, mockModels, "mocked_address")

	// asserions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "failed to request external resource", err.Error())

	mockLogger.AssertCalled(t, "Errorf", "failed to request address or utxo", mock.Anything)
	mockModels.AssertExpectations(t)
}
