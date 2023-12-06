package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Tx(context *gin.Context) {
	tx := context.Param("tx")
	transaction, err := services.Tx(models, tx)
	if err != nil {
		sendError(context, http.StatusBadGateway, err.Error())
	}

	sendSuccess(context, *transaction)
}
