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

type Models struct {
	cacher ICacher
}

func NewModels() *Models {
	return &Models{
		cacher: &MemCache{
			data: make(map[string]interface{}),
		},
	}
}

func NewFetcher() *Fetcher {
	return &Fetcher{
		BaseURL:  os.Getenv("BASE_URL"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		logger:   config.GetLogger("model"),
	}
}
