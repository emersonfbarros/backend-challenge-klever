package controller

import (
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendBtcConverted struct {
	Address string
	Amount  *big.Int
}

func Send(context *gin.Context) {
	request := SendBtc{}
	btcTransactionData := SendBtcConverted{}

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

	btcTransactionData.Address = request.Address
	btcTransactionData.Amount = new(big.Int)
	btcTransactionData.Amount.SetString(request.Amount, 10)

	sendSuccess(context, btcTransactionData)
}
