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

var logger *config.Logger

var fetcher model.IFetcher

var services *Services

func InitService() *Services {
	logger = config.GetLogger("service")
	_, fetcher = model.InitModel()

	services = &Services{}
	return services
}
