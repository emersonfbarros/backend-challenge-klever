package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IResSender interface {
	sendError(ctx *gin.Context, code int, errMsg string)
	sendSuccess(ctx *gin.Context, data interface{})
}

type ResSender struct{}

func NewResSender() *ResSender {
	return &ResSender{}
}

func (rs *ResSender) sendError(context *gin.Context, code int, errMsg string) {
	context.Header("Content-type", "application/json")
	context.JSON(code, gin.H{
		"message": errMsg,
	})
}

func (rs *ResSender) sendSuccess(context *gin.Context, data interface{}) {
	context.Header("Content-type", "application/json")
	context.JSON(http.StatusOK, data)
}
