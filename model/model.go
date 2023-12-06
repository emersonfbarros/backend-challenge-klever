package model

import (
	"os"

	"github.com/emersonfbarros/backend-challenge-klever/config"
)

type Fetcher interface {
	Fetch(route string, value string) ([]byte, error)
}

type IModels interface {
	Utxo(fetcher Fetcher, address string) (*[]UtxoConverted, error)
	Address(fetcher Fetcher, address string) (*AddressRes, error)
	GetTx(fetcher Fetcher, txId string) (*ExtTx, error)
}

type APIClient struct {
	BaseURL  string
	Username string
	Password string
}

type Models struct{}

var logger *config.Logger

func NewModels() *Models {
	return &Models{}
}

// instantiates Models and APIClient structs
// to use the model layer methods
func InitModel() (*Models, *APIClient) {
	logger = config.GetLogger("model")
	return NewModels(), &APIClient{
		BaseURL:  os.Getenv("BASE_URL"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
}
