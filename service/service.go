package service

import (
	"github.com/emersonfbarros/backend-challenge-klever/config"
	"github.com/emersonfbarros/backend-challenge-klever/model"
)

type IServices interface {
	BalanceCalc(models model.IModels, address string) (*BalanceResult, error, int)
	Details(services IServices, models model.IModels, address string) (*AddressInfo, error, int)
	Tx(models model.IModels, txId string) (*Transaction, error, int)
	Send(models model.IModels, btcTransactionData *SendBtcConverted) (*UtxoNeeded, error, int)
	Health(fetcher model.IFetcher) *HealthRes
}

type Services struct{}

func NewServices() *Services {
	return &Services{}
}

var logger config.ILogger
var fetcher model.IFetcher

func InitService() {
	logger = config.GetLogger("service")
	fetcher = model.NewFetcher()
}
