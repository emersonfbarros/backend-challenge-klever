package controller

import (
	"github.com/gin-gonic/gin"
)

func Tx(context *gin.Context) {
	tx := context.Param("tx")
	transaction, err, httpCode := services.Tx(models, tx)
	if err != nil {
		resSender.sendError(context, httpCode, err.Error())
		return
	}

	resSender.sendSuccess(context, transaction)
}
