package service

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/emersonfbarros/backend-challenge-klever/model"
)

type SendBtcConverted struct {
	Address string
	Amount  *big.Int
}

type Utxo struct {
	Txid   string `json:"txid"`
	Amount string `json:"amount"`
}

type UtxoNeeded struct {
	Utxos []Utxo `json:"utxos"`
}

func (s *Services) Send(models model.IModels, btcTransactionData *SendBtcConverted) (*UtxoNeeded, error, int) {
	utxoRef, err, httpCode := models.Utxo(fetcher, btcTransactionData.Address)
	if err != nil {
		logger.Errorf("%s", err.Error())
		return nil, fmt.Errorf("%s", err), httpCode
	}

	// sort utxo slice in descending order
	utxoSlice := *utxoRef
	sort.Slice(utxoSlice, func(i, j int) bool {
		return utxoSlice[i].Value.Cmp(utxoSlice[j].Value) > 0
	})

	var utxosNeeded UtxoNeeded
	compValue := big.NewInt(0) // comparision value
	for _, utxo := range utxoSlice {
		compValue.Add(compValue, utxo.Value) // add utxo.Value to compValue

		// fill utxoNeeded.Utxo
		utxosNeeded.Utxos = append(utxosNeeded.Utxos, Utxo{
			Txid:   utxo.Txid,
			Amount: utxo.Value.String(),
		})

		// if compValue is greater than btcTransactionData.Amount break the loop
		if compValue.Cmp(btcTransactionData.Amount) > 0 {
			break
		}
	}

	return &utxosNeeded, nil, 0
}
