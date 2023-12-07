package service

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockFetcher é uma implementação mock da interface IFetcher para uso nos testes.
type MockFetcher struct {
	mock.Mock
}

func (m *MockFetcher) Fetch(route string, value string) ([]byte, error) {
	args := m.Called(route, value)
	return args.Get(0).([]byte), args.Error(1)
}

func TestHealthSuccess(t *testing.T) {
	// config env
	os.Setenv("ADDRESS_TEST", "mock_address")
	os.Setenv("TX_TEST", "mock_tx")

	mockFetcher := new(MockFetcher)

	// config mocks
	mockFetcher.On("Fetch", "address", "mock_address").Return([]byte("address_data"), nil)
	mockFetcher.On("Fetch", "utxo", "mock_address").Return([]byte("utxo_data"), nil)
	mockFetcher.On("Fetch", "tx", "mock_tx").Return([]byte("tx_data"), nil)

	// creates service and call Health method
	service := &Services{}
	result := service.Health(mockFetcher)

	// asserts
	assert.NotNil(t, result)
	assert.Equal(t, "healthy", result.Status)
	assert.NotEmpty(t, result.Timestamp)
	assert.NotEmpty(t, result.Uptime)
	assert.NotEmpty(t, result.ExternalApi.Status)
	assert.NotEmpty(t, result.ExternalApi.ResponseTime)

	mockFetcher.AssertCalled(t, "Fetch", "address", "mock_address")
	mockFetcher.AssertCalled(t, "Fetch", "utxo", "mock_address")
	mockFetcher.AssertCalled(t, "Fetch", "tx", "mock_tx")

	// clean env configs
	os.Unsetenv("ADDRESS_TEST")
	os.Unsetenv("TX_TEST")
}

func TestHealthPartialSuccess(t *testing.T) {
	// config env
	os.Setenv("ADDRESS_TEST", "mock_address")
	os.Setenv("TX_TEST", "mock_tx")

	mockFetcher := new(MockFetcher)

	// config mocks
	mockFetcher.On("Fetch", "address", "mock_address").Return([]byte("nil"), errors.New("failed to fetch address route"))
	mockFetcher.On("Fetch", "utxo", "mock_address").Return([]byte("utxo_data"), nil)
	mockFetcher.On("Fetch", "tx", "mock_tx").Return([]byte("tx_data"), nil)

	// creates service and call Health method
	service := &Services{}
	result := service.Health(mockFetcher)

	// asserts
	assert.NotNil(t, result)
	assert.Equal(t, "healthy", result.Status)
	assert.NotEmpty(t, result.Timestamp)
	assert.NotEmpty(t, result.Uptime)
	assert.NotEmpty(t, result.ExternalApi.Status)
	assert.NotEmpty(t, result.ExternalApi.ResponseTime)

	mockFetcher.AssertCalled(t, "Fetch", "address", "mock_address")
	mockFetcher.AssertCalled(t, "Fetch", "utxo", "mock_address")
	mockFetcher.AssertCalled(t, "Fetch", "tx", "mock_tx")

	// clean env configs
	os.Unsetenv("ADDRESS_TEST")
	os.Unsetenv("TX_TEST")
}

func TestHealthExternalApiDown(t *testing.T) {
	// config env
	os.Setenv("ADDRESS_TEST", "mock_address")
	os.Setenv("TX_TEST", "mock_tx")

	mockFetcher := new(MockFetcher)

	// config mocks
	mockFetcher.On("Fetch", "address", "mock_address").Return([]byte("nil"), errors.New("failed to fetch address route"))
	mockFetcher.On("Fetch", "utxo", "mock_address").Return([]byte("nil"), errors.New("failed to fetch utxo route"))
	mockFetcher.On("Fetch", "tx", "mock_tx").Return([]byte("nil"), errors.New("failed to fetch tx route"))

	// creates service and call Health method
	service := &Services{}
	result := service.Health(mockFetcher)

	// asserts
	assert.NotNil(t, result)
	assert.Equal(t, "healthy", result.Status)
	assert.NotEmpty(t, result.Timestamp)
	assert.NotEmpty(t, result.Uptime)
	assert.NotEmpty(t, result.ExternalApi.Status)
	assert.NotEmpty(t, result.ExternalApi.ResponseTime)

	mockFetcher.AssertCalled(t, "Fetch", "address", "mock_address")
	mockFetcher.AssertCalled(t, "Fetch", "utxo", "mock_address")
	mockFetcher.AssertCalled(t, "Fetch", "tx", "mock_tx")

	// clean env configs
	os.Unsetenv("ADDRESS_TEST")
	os.Unsetenv("TX_TEST")
}
