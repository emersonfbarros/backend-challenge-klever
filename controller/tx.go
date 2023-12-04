package controller

import "github.com/gin-gonic/gin"

func Tx(context *gin.Context) {
	tx := context.Param("tx")
	sendSuccess(context, tx)
}
