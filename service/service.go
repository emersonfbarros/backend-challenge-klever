package service

import (
	"github.com/emersonfbarros/backend-challenge-klever/config"
	"github.com/emersonfbarros/backend-challenge-klever/model"
)

type IServices interface {
	BalanceCalc(models model.IModels, address string) (*BalanceResult, error)
	Details(models model.IModels, address string) (*AddressInfo, error)
	Tx(models model.IModels, txId string) (*Transaction, error)
	Send(models model.IModels, btcTransactionData *SendBtcConverted) (*UtxoNeeded, error)
	Health(fetcher model.IFetcher) *HealthRes
}

type Services struct{}

func NewServices() *Services {
	return &Services{}
}

var logger config.ILogger
var fetcher model.IFetcher
var services *Services

func InitService() {
	logger = config.GetLogger("service")
	model.InitModel()
	fetcher = model.NewFetcher()
	services = NewServices()
}
