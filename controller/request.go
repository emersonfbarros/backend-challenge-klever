package controller

import "fmt"

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
	return nil
}
