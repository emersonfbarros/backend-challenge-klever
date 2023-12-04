package controller

import (
	"github.com/emersonfbarros/backend-challenge-klever/service"
	"github.com/gin-gonic/gin"
)

func Balance(context *gin.Context) {
	address := context.Param("address")
	service.Balance(address)
	sendSuccess(context, address)
}
