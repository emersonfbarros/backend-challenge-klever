package controller

import (
	"github.com/gin-gonic/gin"
)

func Balance(context *gin.Context) {
	address := context.Param("address")

	balance, err, httpCode := services.BalanceCalc(models, address)
	if err != nil {
		resSender.sendError(context, httpCode, err.Error())
		return
	}

	resSender.sendSuccess(context, balance)
}
