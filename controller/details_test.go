package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emersonfbarros/backend-challenge-klever/service"
	"github.com/gin-gonic/gin"
)

func TestDetailsSuccess(t *testing.T) {
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

	detailsResult := &service.AddressInfo{
		Address:         "valid_address",
		BalanceExternal: "38675889",
		TotalTx:         31,
		Balance: service.BalanceResult{
			Confirmed:   "38675889",
			Unconfirmed: "8527",
		},
		Total: service.Total{
			Sent:     "38675889",
			Received: "38675889",
		},
	}

	s.On("Details", s, m, "valid_address").Return(detailsResult, nil).Once()

	r.On("sendSuccess", context, detailsResult).Once()

	Details(context)

	s.AssertExpectations(t)
	r.AssertExpectations(t)
}

func TestDetailsError(t *testing.T) {
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

	detailsResult := &service.AddressInfo{
		Address:         "nil",
		BalanceExternal: "nil",
		TotalTx:         0,
		Balance: service.BalanceResult{
			Confirmed:   "nil",
			Unconfirmed: "nil",
		},
		Total: service.Total{
			Sent:     "nil",
			Received: "nil",
		},
	}

	s.On("Details", s, m, "valid_address").Return(detailsResult, errors.New("failed to get details")).Once()

	r.On("sendError", context, http.StatusBadGateway, "failed to get details").Once()

	Details(context)

	s.AssertExpectations(t)
	r.AssertExpectations(t)
}
