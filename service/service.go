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
}

type Services struct{}

var logger *config.Logger

var models model.IModels
var fetcher model.Fetcher

var services *Services

func InitService() *Services {
	logger = config.GetLogger("service")
	models, fetcher = model.InitModel()

	services = &Services{}
	return services
}
