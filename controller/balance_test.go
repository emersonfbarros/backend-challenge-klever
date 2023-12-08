package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emersonfbarros/backend-challenge-klever/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
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

	balanceResult := &service.BalanceResult{
		Confirmed:   "38675889",
		Unconfirmed: "8527",
	}

	s.On("BalanceCalc", m, mock.Anything).Return(balanceResult, nil).Once()

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

	balanceResult := &service.BalanceResult{
		Confirmed:   "nil",
		Unconfirmed: "nil",
	}

	s.On("BalanceCalc", m, mock.Anything).Return(balanceResult, errors.New("failed to get balance")).Once()

	r.On("sendError", context, http.StatusBadGateway, "failed to get balance").Once()

	Balance(context)

	s.AssertExpectations(t)
	r.AssertExpectations(t)
}
