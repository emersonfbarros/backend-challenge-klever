package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emersonfbarros/backend-challenge-klever/service"
	"github.com/gin-gonic/gin"
)

func TestBalanceSuccess(t *testing.T) {
	InitController()
	s := new(MockServices) // implemented in controler/tx_test
	services = s
	m := new(MockIModels) // implemented in controler/tx_test
	models = m
	r := new(MockResSender) // implemented in controler/tx_test
	resSender = r

	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	context.Params = append(context.Params, gin.Param{Key: "address", Value: "valid_address"})

	balanceResult := &service.BalanceResult{
		Confirmed:   "38675889",
		Unconfirmed: "8527",
	}

	s.On("BalanceCalc", m, "valid_address").Return(balanceResult, nil, http.StatusOK).Once()

	r.On("sendSuccess", context, balanceResult).Once()

	Balance(context)

	s.AssertExpectations(t)
	r.AssertExpectations(t)
}

func TestBalanceError(t *testing.T) {
	InitController()
	s := new(MockServices) // implemented in controler/tx_test
	services = s
	m := new(MockIModels) // implemented in controler/tx_test
	models = m
	r := new(MockResSender) // implemented in controler/tx_test
	resSender = r

	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	context.Params = append(context.Params, gin.Param{Key: "address", Value: "valid_address"})

	balanceResult := &service.BalanceResult{
		Confirmed:   "nil",
		Unconfirmed: "nil",
	}

	s.On("BalanceCalc", m, "valid_address").Return(balanceResult, errors.New("failed to get balance"), http.StatusBadGateway).Once()


	r.On("sendError", context, http.StatusBadGateway, "failed to get balance").Once()

	Balance(context)

	s.AssertExpectations(t)
	r.AssertExpectations(t)
}
