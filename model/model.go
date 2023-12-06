package model

import (
	"os"

	"github.com/emersonfbarros/backend-challenge-klever/config"
)

type Fetcher interface {
	Fetch(route string, value string) ([]byte, error)
}

type DataProcessors interface {
	Utxo(fetcher Fetcher, address string) (*[]UtxoConverted, error)
	Address(fetcher Fetcher, address string) (*AddressRes, error)
	GetTx(fetcher Fetcher, txId string) (*ExtTx, error)
}

type APIClient struct {
	BaseURL  string
	Username string
	Password string
}

type APIDataHandler struct{}

var logger *config.Logger

// instantiates APIDataHandler and APIClient structs
// to use the model layer methods
func InitModel() (*APIDataHandler, *APIClient) {
	logger = config.GetLogger("model")
	return &APIDataHandler{}, &APIClient{
		BaseURL:  os.Getenv("BASE_URL"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
}
