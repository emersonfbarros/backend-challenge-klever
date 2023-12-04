package controller

import (
	"fmt"
	"math/big"
)

type SentBtc struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

func (r *SentBtc) Validate() error {
	if r.Address == "" && r.Amount == "" {
		return fmt.Errorf("Request body is empty or malformed")
	}
	if r.Address == "" {
		return fmt.Errorf("'address' is required")
	}
	if r.Amount == "" {
		return fmt.Errorf("'amount' is required")
	}

	a := new(big.Int)
	_, ok := a.SetString(r.Amount, 10)
	if !ok {
		return fmt.Errorf("'amount' must be a valid number")
	}
	return nil
}
