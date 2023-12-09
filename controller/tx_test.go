package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emersonfbarros/backend-challenge-klever/model"
	"github.com/emersonfbarros/backend-challenge-klever/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// mock IServices
type MockServices struct {
	mock.Mock
}

func (m *MockServices) BalanceCalc(models model.IModels, address string) (*service.BalanceResult, error, int) {
	args := m.Called(models, address)
	return args.Get(0).(*service.BalanceResult), args.Error(1), args.Int(2)
}

func (m *MockServices) Details(services service.IServices, models model.IModels, address string) (*service.AddressInfo, error, int) {
	args := m.Called(services, models, address)
	return args.Get(0).(*service.AddressInfo), args.Error(1), args.Int(2)
}

func (m *MockServices) Tx(models model.IModels, txId string) (*service.Transaction, error, int) {
	args := m.Called(models, txId)
	return args.Get(0).(*service.Transaction), args.Error(1), args.Int(2)
}

func (m *MockServices) Send(models model.IModels, btcTransactionData *service.SendBtcConverted) (*service.UtxoNeeded, error, int) {
	args := m.Called(models, btcTransactionData)
	return args.Get(0).(*service.UtxoNeeded), args.Error(1), args.Int(2)
}

func (m *MockServices) Health(fetcher model.IFetcher) *service.HealthRes {
	args := m.Called(fetcher)
	return args.Get(0).(*service.HealthRes)
}

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

// MockResSender mocks IResSender.
type MockResSender struct {
	mock.Mock
}

func (m *MockResSender) sendError(ctx *gin.Context, code int, errMsg string) {
	m.Called(ctx, code, errMsg)
}

func (m *MockResSender) sendSuccess(ctx *gin.Context, data interface{}) {
	m.Called(ctx, data)
}

func TestTxSuccess(t *testing.T) {
	InitController()
	s := new(MockServices)
	services = s
	m := new(MockIModels)
	models = m
	r := new(MockResSender)
	resSender = r
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)

	transaction := &service.Transaction{
		TxID: "hdfghqewihgiqhgqw",
		Addresses: []service.Address{
			{Address: "h89ef71h7yqgghbqjiufvsdy8967", Value: "100"},
		},
		Block: 974235,
	}

	s.On("Tx", m, mock.Anything).Return(transaction, nil, http.StatusOK).Once()

	r.On("sendSuccess", context, transaction).Once()

	Tx(context)

	s.AssertExpectations(t)
	r.AssertExpectations(t)
}

func TestTxError(t *testing.T) {
	InitController()
	s := new(MockServices)
	services = s
	m := new(MockIModels)
	models = m
	r := new(MockResSender)
	resSender = r
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)

	transaction := &service.Transaction{
		TxID: "nil",
		Addresses: []service.Address{
			{Address: "nil", Value: "nil"},
		},
		Block: 0,
	}

	s.On("Tx", m, mock.Anything).Return(transaction, errors.New("failed to get transaction"), http.StatusBadGateway).Once()

	r.On("sendError", context, http.StatusBadGateway, "failed to get transaction").Once()

	Tx(context)

	s.AssertExpectations(t)
	r.AssertExpectations(t)
}
