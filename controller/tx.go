package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Tx(context *gin.Context) {
	tx := context.Param("tx")
	transaction, err := services.Tx(models, tx)
	if err != nil {
		resSender.sendError(context, http.StatusBadGateway, err.Error())
	}

	resSender.sendSuccess(context, *transaction)
}
