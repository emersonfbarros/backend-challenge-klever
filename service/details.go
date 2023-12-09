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

func (s *Services) Details(services IServices, models model.IModels, address string) (*AddressInfo, error, int) {
	var wg sync.WaitGroup
	wg.Add(2)

	var detailsRef *model.AddressRes
	var balanceRef *BalanceResult
	var errDt, errBl error
	var httpCode int

	// fetch external api simultaneously
	go func() {
		defer wg.Done()
		balanceRef, errBl, httpCode = services.BalanceCalc(models, address)
	}()

	go func() {
		defer wg.Done()
		detailsRef, errDt, httpCode = models.Address(fetcher, address)
	}()

	// wait for both requests to complete
	wg.Wait()

	if errBl != nil {
		return nil, errBl, httpCode
	}

	if errDt != nil {
		return nil, errDt, httpCode
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

	return &adressInfo, nil, httpCode
}
