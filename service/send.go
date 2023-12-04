package service

import (
	"math/big"
)

type SendBtcConverted struct {
	Address string
	Amount  *big.Int
}

func Send(btcTransactionData *SendBtcConverted) {
	logger.Infof("RECEIVED TRANSACTION: %+v", btcTransactionData)
}
