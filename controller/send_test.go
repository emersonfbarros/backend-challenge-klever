package controller

import (
	"bytes"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emersonfbarros/backend-challenge-klever/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSendSuccess(t *testing.T) {
	// setup mocks
	InitController()
	mockServices := new(MockServices) // implemented in controler/tx_test
	services = mockServices
	mockModels := new(MockIModels) // implemented in controler/tx_test
	models = mockModels
	mockResSend := new(MockResSender) // implemented in controler/tx_test
	resSender = mockResSend

	// setup final result
	expectedUtxo := &service.UtxoNeeded{
		Utxos: []service.Utxo{
			{
				Txid:   "3672b48dcb1494a30ad94b2cd092cc380d6bb475b86ca547655a65c0c27941e5",
				Amount: "1146600",
			},
			{
				Txid:   "4148e2e46000c3a988ae2b7687a40b2bab7f29fb822ecc22912566d7b74330a4",
				Amount: "1100593",
			},
		},
	}

	// setup req data
	requestBody := sendBtc{
		Address: "valid_address",
		Amount:  "1208053",
	}

	// setup request
	requestBodyBytes, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/send", bytes.NewReader(requestBodyBytes))

	// creates false gin context
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	context.Request = req

	// setup param for services.Send
	expectedDataToCallSend := service.SendBtcConverted{
		Address: "valid_address",
		Amount:  big.NewInt(1208053),
	}

	// config mocks
	mockServices.On("Send", models, &expectedDataToCallSend).Return(expectedUtxo, nil, 0)

	mockResSend.On("sendSuccess", context, expectedUtxo)

	Send(context)

	// assertions
	assert.Equal(t, http.StatusOK, w.Code)

	mockServices.AssertExpectations(t)
	mockResSend.AssertExpectations(t)
}

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Infof(format string, args ...interface{}) {
	m.Called(format, args)
}

func (m *MockLogger) Errorf(format string, args ...interface{}) {
	m.Called(format, args)
}

func TestSendValidationError(t *testing.T) {
	cases := []struct {
		address  string
		amount   string
		testName string
		errMsg   string
	}{
		{
			testName: "malformed",
			address:  "",
			amount:   "",
			errMsg:   "Request body is empty or malformed",
		},
		{
			testName: "missing address",
			address:  "",
			amount:   "1208053",
			errMsg:   "'address' is required",
		},
		{
			testName: "missing amount",
			address:  "valid_address",
			amount:   "",
			errMsg:   "'amount' is required",
		},
		{
			testName: "invalid amount",
			address:  "valid_address",
			amount:   "500O",
			errMsg:   "'amount' must be a valid number",
		},
		{
			testName: "amount zero",
			address:  "valid_address",
			amount:   "0",
			errMsg:   "'amount' must be greater than zero",
		},
	}

	for _, tt := range cases {
		t.Run(tt.testName, func(t *testing.T) {
			// setup mocks
			InitController()
			mockLogger := new(MockLogger)
			logger = mockLogger
			mockResSend := new(MockResSender) // implemented in controler/tx_test
			resSender = mockResSend

			reqBody := sendBtc{
				Address: tt.address,
				Amount:  tt.amount,
			}

			// setup request
			requestBodyBytes, _ := json.Marshal(reqBody)
			req, _ := http.NewRequest("POST", "/send", bytes.NewReader(requestBodyBytes))

			// creates false gin context
			w := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(w)
			context.Request = req

			// config mocks
			mockResSend.On("sendError", context, http.StatusBadRequest, tt.errMsg)
			mockLogger.On("Errorf", "validation error: %v", mock.Anything)

			Send(context)

			mockResSend.AssertExpectations(t)
			mockLogger.AssertExpectations(t)
		})
	}
}
