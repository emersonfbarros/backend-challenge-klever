package controller

import (
	"math/big"
	"net/http"

	"github.com/emersonfbarros/backend-challenge-klever/service"
	"github.com/gin-gonic/gin"
)

func Send(context *gin.Context) {
	request := sendBtc{}
	btcTransactionData := service.SendBtcConverted{}

	if err := context.BindJSON(&request); err != nil {
		logger.Errorf("json bind error: %v", err)
		resSender.sendError(context, http.StatusBadRequest, "'address' and 'amount' must be strings")
		return
	}

	// validate request body
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		resSender.sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	btcTransactionData.Address = request.Address
	// convert amount to big.Int
	btcTransactionData.Amount = new(big.Int)
	btcTransactionData.Amount.SetString(request.Amount, 10)

	utxos, err, httpCode := services.Send(models, &btcTransactionData)
	if err != nil {
		resSender.sendError(context, httpCode, err.Error())
		return
	}

	resSender.sendSuccess(context, utxos)
}
