package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Details(context *gin.Context) {
	address := context.Param("address")
	details, err := services.Details(models, address)
	if err != nil {
		sendError(context, http.StatusBadGateway, err.Error())
	}

	sendSuccess(context, *details)
}
