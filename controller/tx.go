package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Tx(context *gin.Context) {
	tx := context.Param("tx")
	transaction, err := services.Tx(models, tx)
	if err != nil {
		fmt.Printf("\n\nENTROU\n\n")
		resSender.sendError(context, http.StatusBadGateway, err.Error())
		return
	}

	resSender.sendSuccess(context, transaction)
}
