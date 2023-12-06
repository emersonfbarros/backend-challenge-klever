package model

import (
	"os"

	"github.com/emersonfbarros/backend-challenge-klever/config"
)

type IFetcher interface {
	Fetch(route string, value string) ([]byte, error)
}

type IModels interface {
	Utxo(fetcher IFetcher, address string) (*[]UtxoConverted, error)
	Address(fetcher IFetcher, address string) (*AddressRes, error)
	GetTx(fetcher IFetcher, txId string) (*ExtTx, error)
}

type Fetcher struct {
	BaseURL  string
	Username string
	Password string
}

type Models struct{}

func NewModels() *Models {
	return &Models{}
}

func NewFetcher() *Fetcher {
	return &Fetcher{
		BaseURL:  os.Getenv("BASE_URL"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
}

var logger config.ILogger

func InitModel() {
	logger = config.GetLogger("model")
}
