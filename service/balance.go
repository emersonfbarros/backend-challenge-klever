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

func (service *Services) BalanceCalc(models model.IModels, address string) (*BalanceResult, error) {
	utxoRef, err := models.Utxo(fetcher, address)
	if err != nil {
		logger.Errorf("failed to unmarshal api response %v", err.Error())
		return nil, fmt.Errorf("failed to request external resouce")
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

	return &balanceResult, nil
}
