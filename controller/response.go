package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(context *gin.Context, code int, msg string) {
	context.Header("Content-type", "application/json")
	context.JSON(code, gin.H{
		"message": msg,
	})
}

func sendSuccess(context *gin.Context, data interface{}) {
	context.Header("Content-type", "application/json")
	context.JSON(http.StatusOK, data)
}
