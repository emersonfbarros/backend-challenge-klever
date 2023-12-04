package controller

import (
	"github.com/emersonfbarros/backend-challenge-klever/service"
	"github.com/gin-gonic/gin"
)

func Details(context *gin.Context) {
	address := context.Param("address")
	service.Details(address)
	sendSuccess(context, address)
}
