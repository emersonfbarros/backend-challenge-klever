package controller

import "github.com/gin-gonic/gin"

func Details(context *gin.Context) {
	address := context.Param("address")
	sendSuccess(context, address)
}
