package controller

import (
	"net/http"

	"github.com/emersonfbarros/backend-challenge-klever/service"
	"github.com/gin-gonic/gin"
)

func Details(context *gin.Context) {
	address := context.Param("address")
	details, err := service.Details(address)
	if err != nil {
		sendError(context, http.StatusBadGateway, err.Error())
	}

	sendSuccess(context, *details)
}
