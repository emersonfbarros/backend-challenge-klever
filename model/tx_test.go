package model

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTxSuccess(t *testing.T) {
	expectedTx := ExtTx{
		Vin:         []vin{{Addresses: []string{"address1"}, Value: "100"}},
		Vout:        []vout{{Value: "100", Addresses: []string{"address2"}}},
		BlockHeight: 123,
	}
	expectedTxBytes, _ := json.Marshal(expectedTx)

	// initialize mock, declared in model/address_test
	mockFetcher := new(MockFetcher)

	mockFetcher.On("Fetch", "tx", "tx123").Return(expectedTxBytes, nil)

	models := &Models{}

	resultTx, err := models.GetTx(mockFetcher, "tx123")

	// Asserts
	assert.NoError(t, err)
	assert.Equal(t, &expectedTx, resultTx)

	mockFetcher.AssertExpectations(t)
}

func TestGetTxErrorFetch(t *testing.T) {

	expectedTxBytes, _ := json.Marshal(nil)

	// initialize mock, declared in model/address_test
	mockFetcher := new(MockFetcher)

	mockFetcher.On("Fetch", "tx", "tx123").Return(expectedTxBytes, errors.New("fetch error"))

	models := &Models{}

	resultTx, err := models.GetTx(mockFetcher, "tx123")

	// Asserts
	assert.Error(t, err)
	assert.Nil(t, resultTx)

	mockFetcher.AssertExpectations(t)
}

func TestGetTxErrorUnmarshal(t *testing.T) {

	expectedTxBytes := []byte("invalid")

	// initialize mock, declared in model/address_test
	mockFetcher := new(MockFetcher)

	mockFetcher.On("Fetch", "tx", "tx123").Return(expectedTxBytes, errors.New("fetch error"))

	models := &Models{}

	resultTx, err := models.GetTx(mockFetcher, "tx123")

	// Asserts
	assert.Error(t, err)
	assert.Nil(t, resultTx)

	mockFetcher.AssertExpectations(t)
}
