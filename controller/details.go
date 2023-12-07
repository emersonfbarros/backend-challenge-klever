package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Details(context *gin.Context) {
	address := context.Param("address")
	details, err := services.Details(services, models, address)
	if err != nil {
		resSender.sendError(context, http.StatusBadGateway, err.Error())
	}

	resSender.sendSuccess(context, *details)
}
