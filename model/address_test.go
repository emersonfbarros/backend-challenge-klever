package model

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mock IFetcher
type MockFetcher struct {
	mock.Mock
}

func (m *MockFetcher) Fetch(route string, value string) ([]byte, error) {
	args := m.Called(route, value)
	return args.Get(0).([]byte), args.Error(1)
}

func TestAddressSuccess(t *testing.T) {
	// creates fetcher mock and model
	mockFetcher := new(MockFetcher)

	models := &Models{}

	// test data
	testAddress := "test_address"
	testResponse := &AddressRes{
		Balance:       "1000",
		Txs:           2,
		TotalSent:     "500",
		TotalReceived: "1500",
	}
	expectedBytes, _ := json.Marshal(testResponse)

	mockFetcher.On("Fetch", "address", testAddress).Return(expectedBytes, nil)

	// calls method
	result, err, httpCode := models.Address(mockFetcher, testAddress)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, testResponse, result)
	assert.Equal(t, 0, httpCode)

	mockFetcher.AssertExpectations(t)
}

func TestAddressErrorFetch(t *testing.T) {
	expectedBytes, _ := json.Marshal(nil)

	testAddress := "test_address"

	mockFetcher := new(MockFetcher)

	models := &Models{}

	mockFetcher.On("Fetch", "address", testAddress).Return(expectedBytes, errors.New("fetch error"))

	// calls method
	result, err, httpCode := models.Address(mockFetcher, testAddress)

	// assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, http.StatusBadGateway, httpCode)

	mockFetcher.AssertExpectations(t)
}

func TestAddressErrorUnmarshal(t *testing.T) {
	// invalid json to unmarshal
	expectedBytes := []byte("invalid")

	testAddress := "test_address"

	mockFetcher := new(MockFetcher)

	models := &Models{}

	mockFetcher.On("Fetch", "address", testAddress).Return(expectedBytes, nil)

	// calls method
	result, err, httpCode := models.Address(mockFetcher, testAddress)

	// assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, http.StatusInternalServerError, httpCode)

	mockFetcher.AssertExpectations(t)
}
