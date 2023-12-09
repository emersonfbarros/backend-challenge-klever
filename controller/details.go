package controller

import "github.com/gin-gonic/gin"

func Details(context *gin.Context) {
	address := context.Param("address")
	details, err, httpCode := services.Details(services, models, address)
	if httpCode != 0 {
		resSender.sendError(context, httpCode, err.Error())
		return
	}

	resSender.sendSuccess(context, details)
}
