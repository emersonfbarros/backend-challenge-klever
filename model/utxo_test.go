package model

import (
	"errors"
	"math/big"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtxoSuccess(t *testing.T) {
	expectedTxBytes := []byte(`[{"txid":"123","value":"456","confirmations":3}]`)

	testAddress := "test_address"

	// initialize mock, declared in model/address_test
	mockFetcher := new(MockFetcher)

	mockFetcher.On("Fetch", "utxo", testAddress).Return(expectedTxBytes, nil)

	models := Models{}

	result, err, httpCode := models.Utxo(mockFetcher, testAddress)

	// assertions
	assert.Nil(t, err)

	assert.Len(t, *result, 1)
	assert.Equal(t, "123", (*result)[0].Txid)
	assert.Equal(t, big.NewInt(456), (*result)[0].Value)
	assert.Equal(t, 3, (*result)[0].Confirmations)
	assert.Equal(t, 0, httpCode)

	mockFetcher.AssertExpectations(t)
}

func TestUtxoErrorFetch(t *testing.T) {
	expectedTxBytes := []byte(nil)

	testAddress := "test_address"

	// initialize mock, declared in model/address_test
	mockFetcher := new(MockFetcher)

	mockFetcher.On("Fetch", "utxo", testAddress).Return(expectedTxBytes, errors.New("fetch error"))

	models := Models{}

	result, err, httpCode := models.Utxo(mockFetcher, testAddress)

	// assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, http.StatusBadGateway, httpCode)

	mockFetcher.AssertExpectations(t)
}

func TestUtxoErrorUnmarshal(t *testing.T) {
	// invalid json to unmarshal
	expectedTxBytes := []byte("invalid")

	testAddress := "test_address"

	// initialize mock, declared in model/address_test
	mockFetcher := new(MockFetcher)

	mockFetcher.On("Fetch", "utxo", testAddress).Return(expectedTxBytes, nil)

	models := Models{}

	result, err, httpCode := models.Utxo(mockFetcher, testAddress)

	// assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, http.StatusInternalServerError, httpCode)

	mockFetcher.AssertExpectations(t)
}
