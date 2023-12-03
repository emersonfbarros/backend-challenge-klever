package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Send(context *gin.Context) {
	request := SentBtc{}

	if err := context.BindJSON(&request); err != nil {
		logger.Errorf("json bind error: %v", err)
		sendError(context, http.StatusBadRequest, "'address' and 'amount' must be strings")
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	sendSuccess(context, request)
}
