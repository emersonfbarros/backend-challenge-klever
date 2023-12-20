package model

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMemCache struct {
	mock.Mock
}

func (mc *MockMemCache) Set(key string, value interface{}) {
	mc.Called(key, value)
}

func (mc *MockMemCache) Get(key string) (interface{}, bool) {
	args := mc.Called(key)
	return args.Get(0), args.Bool(1)
}

func TestGetTxSuccessWithoutCache(t *testing.T) {
	expectedTx := ExtTx{
		Vin:         []Vin{{Addresses: []string{"address1"}, Value: "100"}},
		Vout:        []Vout{{Value: "100", Addresses: []string{"address2"}}},
		BlockHeight: 123,
	}
	expectedTxBytes, _ := json.Marshal(expectedTx)

	// initialize mock, declared in model/address_test
	mockFetcher := new(MockFetcher)
	mockCacher := new(MockMemCache)

	mockFetcher.On("Fetch", "tx", "tx123").Return(expectedTxBytes, nil)
	mockCacher.On("Get", "tx123").Return(nil, false)
	mockCacher.On("Set", "tx123", expectedTx).Return()

	models := &Models{
		cacher: mockCacher,
	}

	resultTx, err, httpCode := models.GetTx(mockFetcher, "tx123")

	// Asserts
	assert.NoError(t, err)
	assert.Equal(t, &expectedTx, resultTx)
	assert.Equal(t, 0, httpCode)

	mockFetcher.AssertExpectations(t)
	mockCacher.AssertExpectations(t)
}

func TestGetTxSuccessWithCache(t *testing.T) {
	expectedTx := ExtTx{
		Vin:         []Vin{{Addresses: []string{"address1"}, Value: "100"}},
		Vout:        []Vout{{Value: "100", Addresses: []string{"address2"}}},
		BlockHeight: 123,
	}

	// initialize mock, declared in model/address_test
	mockFetcher := new(MockFetcher)
	mockCacher := new(MockMemCache)

	mockCacher.On("Get", "tx123").Return(expectedTx, true)

	models := &Models{
		cacher: mockCacher,
	}

	resultTx, err, httpCode := models.GetTx(mockFetcher, "tx123")

	// Asserts
	assert.NoError(t, err)
	assert.Equal(t, &expectedTx, resultTx)
	assert.Equal(t, 0, httpCode)

	mockCacher.AssertExpectations(t)
}

func TestGetTxErrorFetch(t *testing.T) {
	expectedTxBytes, _ := json.Marshal(nil)

	// initialize mock, declared in model/address_test
	mockFetcher := new(MockFetcher)
	mockCacher := new(MockMemCache)

	mockFetcher.On("Fetch", "tx", "tx123").Return(expectedTxBytes, errors.New("fetch error"))
	mockCacher.On("Get", "tx123").Return(nil, false)

	models := &Models{
		cacher: mockCacher,
	}

	resultTx, err, httpCode := models.GetTx(mockFetcher, "tx123")

	// Asserts
	assert.Error(t, err)
	assert.Nil(t, resultTx)
	assert.Equal(t, http.StatusBadGateway, httpCode)

	mockFetcher.AssertExpectations(t)
	mockCacher.AssertExpectations(t)
}

func TestGetTxErrorUnmarshal(t *testing.T) {
	// invalid json to unmarshal
	expectedTxBytes := []byte("invalid")

	// initialize mock, declared in model/address_test
	mockFetcher := new(MockFetcher)
	mockCacher := new(MockMemCache)

	mockFetcher.On("Fetch", "tx", "tx123").Return(expectedTxBytes, nil)
	mockCacher.On("Get", "tx123").Return(nil, false)

	models := &Models{
		cacher: mockCacher,
	}

	resultTx, err, httpCode := models.GetTx(mockFetcher, "tx123")

	// Asserts
	assert.Error(t, err)
	assert.Nil(t, resultTx)
	assert.Equal(t, http.StatusInternalServerError, httpCode)

	mockFetcher.AssertExpectations(t)
	mockCacher.AssertExpectations(t)
}
