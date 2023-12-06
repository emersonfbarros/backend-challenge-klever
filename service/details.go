package service

import (
	"sync"

	"github.com/emersonfbarros/backend-challenge-klever/model"
)

type Total struct {
	Sent     string `json:"sent"`
	Received string `json:"received"`
}

type AddressInfo struct {
	Address         string        `json:"address"`
	BalanceExternal string        `json:"balance"`
	TotalTx         int           `json:"totalTx"`
	Balance         BalanceResult `json:"balanceCalc"`
	Total           Total         `json:"total"`
}

func (s *Services) Details(models model.IModels, address string) (*AddressInfo, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	var detailsRef *model.AddressRes
	var balanceRef *BalanceResult
	var errDt, errBl error

	// fetch external api simultaneously
	go func() {
		defer wg.Done()
		balanceRef, errBl = services.BalanceCalc(models, address)
		if errBl != nil {
			logger.Errorf("failed to unmarshal api response %v", errBl.Error())
		}
	}()

	go func() {
		defer wg.Done()
		detailsRef, errDt = models.Address(fetcher, address)
		if errDt != nil {
			logger.Errorf("failed to unmarshal api response %v", errDt.Error())
		}
	}()

	// wait for both requests to complete
	wg.Wait()

	if errBl != nil || errDt != nil {
		logger.Errorf("failed to request address or utxo")
	}

	detailsPartial := *detailsRef
	balanceValues := *balanceRef

	adressInfo := AddressInfo{
		Address:         address,
		BalanceExternal: detailsPartial.Balance,
		TotalTx:         detailsPartial.Txs,
		Balance:         balanceValues,
		Total: Total{
			Sent:     detailsPartial.TotalSent,
			Received: detailsPartial.TotalReceived,
		},
	}

	return &adressInfo, nil
}
