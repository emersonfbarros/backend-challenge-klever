package model

import (
	"os"

	"github.com/emersonfbarros/backend-challenge-klever/config"
)

type IFetcher interface {
	Fetch(route string, value string) ([]byte, error)
}

type IModels interface {
	Utxo(fetcher IFetcher, address string) (*[]UtxoConverted, error, int)
	Address(fetcher IFetcher, address string) (*AddressRes, error, int)
	GetTx(fetcher IFetcher, txId string) (*ExtTx, error, int)
}

type Fetcher struct {
	BaseURL  string
	Username string
	Password string
	logger   config.ILogger
}

var cacher ICacher

func getMemCache() ICacher {
	if cacher == nil {
		cacher = &MemCache{
			data: make(map[string]interface{}),
		}
	}
	return cacher
}

type Models struct{}

func NewModels() *Models {
	getMemCache()
	return &Models{}
}

func NewFetcher() *Fetcher {
	return &Fetcher{
		BaseURL:  os.Getenv("BASE_URL"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		logger:   config.GetLogger("model"),
	}
}
