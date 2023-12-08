package controller

import (
	"net/http/httptest"
	"testing"

	"github.com/emersonfbarros/backend-challenge-klever/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func TestHealth(t *testing.T) {
	InitController()
	s := new(MockServices) // implemented in controler/tx_test
	services = s
	r := new(MockResSender) // implemented in controler/tx_test
	resSender = r

	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)

	healResult := &service.HealthRes{
		Status:    "healthy",
		Timestamp: "2023-12-01T00:00:00Z",
		Uptime:    "0D 3H 20M 30S",
		ExternalApi: service.ExtApi{
			Status:       "Ok",
			ResponseTime: "986ms",
		},
	}

	s.On("Health", mock.Anything).Return(healResult, nil).Once()
	r.On("sendSuccess", context, healResult).Once()

	Health(context)

	s.AssertExpectations(t)
	r.AssertExpectations(t)
}
