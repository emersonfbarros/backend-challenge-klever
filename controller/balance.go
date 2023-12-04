package controller

import (
	"net/http"

	"github.com/emersonfbarros/backend-challenge-klever/service"
	"github.com/gin-gonic/gin"
)

func Balance(context *gin.Context) {
	address := context.Param("address")

	balance, err := service.Balance(address)
	if err != nil {
		sendError(context, http.StatusBadGateway, err.Error())
	}

	sendSuccess(context, *balance)
}
