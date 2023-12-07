package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Balance(context *gin.Context) {
	address := context.Param("address")

	balance, err := services.BalanceCalc(models, address)
	if err != nil {
		resSender.sendError(context, http.StatusBadGateway, err.Error())
		return
	}

	resSender.sendSuccess(context, balance)
}
