package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Send(context *gin.Context) {
	request := SentBtc{}

	err := context.BindJSON(&request)

	if err != nil {
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	sendSuccess(context, request)
}
