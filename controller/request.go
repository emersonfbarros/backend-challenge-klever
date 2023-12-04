package controller

import (
	"fmt"
	"math/big"
)

type sendBtc struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

func (r *sendBtc) Validate() error {
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
	a, ok := a.SetString(r.Amount, 10)
	if !ok {
		return fmt.Errorf("'amount' must be a valid number")
	}
	if a.Sign() <= 0 {
		return fmt.Errorf("'amount' must be greater than zero")
	}
	return nil
}
