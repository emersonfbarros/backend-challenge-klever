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

var logger *config.Logger

func NewModels() *Models {
	return &Models{}
}

// instantiates Models and Fetcher structs
// to use the model layer methods
func InitModel() (*Models, *Fetcher) {
	logger = config.GetLogger("model")
	return NewModels(), &Fetcher{
		BaseURL:  os.Getenv("BASE_URL"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
}
