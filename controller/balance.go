package controller

import "github.com/gin-gonic/gin"

func Balance(context *gin.Context) {
	address := context.Param("address")
	sendSuccess(context, address)
}
