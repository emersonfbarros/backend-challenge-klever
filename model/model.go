package model

import (
	"os"

	"github.com/emersonfbarros/backend-challenge-klever/config"
)

type DataProcessors interface {
	Utxo(fetcher Fetcher, address string) (*[]UtxoConverted, error)
	Address(fetcher Fetcher, address string) (*AddressRes, error)
}

type APIDataHandler struct{}

var logger *config.Logger

func InitModel() {
	logger = config.GetLogger("model")
}
