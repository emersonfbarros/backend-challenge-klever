package controller

import (
	"github.com/emersonfbarros/backend-challenge-klever/service"
	"github.com/gin-gonic/gin"
)

func Tx(context *gin.Context) {
	tx := context.Param("tx")
	service.Tx(tx)
	sendSuccess(context, tx)
}
