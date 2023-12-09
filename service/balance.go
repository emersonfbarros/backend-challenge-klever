package service

import (
	"fmt"
	"math/big"

	"github.com/emersonfbarros/backend-challenge-klever/model"
)

type BalanceResult struct {
	Confirmed   string `json:"confirmed"`
	Unconfirmed string `json:"unconfirmed"`
}

func (s *Services) BalanceCalc(models model.IModels, address string) (*BalanceResult, error, int) {
	utxoRef, err, httpCode := models.Utxo(fetcher, address)
	if err != nil {
		logger.Errorf("%s", err.Error())
		return nil, fmt.Errorf("%s", err.Error()), httpCode
	}

	utxoSlice := *utxoRef
	confirmed := big.NewInt(0)
	unconfirmed := big.NewInt(0)
	for _, utxo := range utxoSlice {
		if utxo.Confirmations < 2 {
			unconfirmed.Add(unconfirmed, utxo.Value)
		} else {
			confirmed.Add(confirmed, utxo.Value)
		}
	}

	balanceResult := BalanceResult{
		Confirmed:   confirmed.String(),
		Unconfirmed: unconfirmed.String(),
	}

	return &balanceResult, nil, 0
}
