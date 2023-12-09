package model

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTxSuccess(t *testing.T) {
	expectedTx := ExtTx{
		Vin:         []Vin{{Addresses: []string{"address1"}, Value: "100"}},
		Vout:        []Vout{{Value: "100", Addresses: []string{"address2"}}},
		BlockHeight: 123,
	}
	expectedTxBytes, _ := json.Marshal(expectedTx)

	// initialize mock, declared in model/address_test
	mockFetcher := new(MockFetcher)

	mockFetcher.On("Fetch", "tx", "tx123").Return(expectedTxBytes, nil)

	models := &Models{}

	resultTx, err, httpCode := models.GetTx(mockFetcher, "tx123")

	// Asserts
	assert.NoError(t, err)
	assert.Equal(t, &expectedTx, resultTx)
	assert.Equal(t, 0, httpCode)

	mockFetcher.AssertExpectations(t)
}

func TestGetTxErrorFetch(t *testing.T) {
	expectedTxBytes, _ := json.Marshal(nil)

	// initialize mock, declared in model/address_test
	mockFetcher := new(MockFetcher)

	mockFetcher.On("Fetch", "tx", "tx123").Return(expectedTxBytes, errors.New("fetch error"))

	models := &Models{}

	resultTx, err, httpCode := models.GetTx(mockFetcher, "tx123")

	// Asserts
	assert.Error(t, err)
	assert.Nil(t, resultTx)
	assert.Equal(t, http.StatusBadGateway, httpCode)

	mockFetcher.AssertExpectations(t)
}

func TestGetTxErrorUnmarshal(t *testing.T) {
	// invalid json to unmarshal
	expectedTxBytes := []byte("invalid")

	// initialize mock, declared in model/address_test
	mockFetcher := new(MockFetcher)

	mockFetcher.On("Fetch", "tx", "tx123").Return(expectedTxBytes, nil)

	models := &Models{}

	resultTx, err, httpCode := models.GetTx(mockFetcher, "tx123")

	// Asserts
	assert.Error(t, err)
	assert.Nil(t, resultTx)
	assert.Equal(t, http.StatusInternalServerError, httpCode)

	mockFetcher.AssertExpectations(t)
}
