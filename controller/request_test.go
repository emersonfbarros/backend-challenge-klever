package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendBtc_Validate(t *testing.T) {
	t.Run("Error cases", func(t *testing.T) {
		// request body empty
		r1 := &sendBtc{}
		err := r1.Validate()
		assert.Error(t, err)

		// address empty
		r2 := &sendBtc{
			Amount: "100",
		}
		err = r2.Validate()
		assert.Error(t, err)
		assert.Equal(t, "'address' is required", err.Error())

		// amount empty
		r3 := &sendBtc{
			Address: "h89ef71h7yqgghbqjiufvsdy8967",
		}
		err = r3.Validate()
		assert.Error(t, err)
		assert.Equal(t, "'amount' is required", err.Error())

		// amount is invalid number
		r4 := &sendBtc{
			Address: "vhfsd8h671-0475yh71hq02f7sfb0",
			Amount:  "abc",
		}
		err = r4.Validate()
		assert.Error(t, err)
		assert.Equal(t, "'amount' must be a valid number", err.Error())

		// amount is less than or equal to zero
		r5 := &sendBtc{
			Address: "cnqe978qbea06tbg1r87ghnjky9",
			Amount:  "0",
		}
		err = r5.Validate()
		assert.Error(t, err)
		assert.Equal(t, "'amount' must be greater than zero", err.Error())
	})

	t.Run("Success cases", func(t *testing.T) {
		// success validation
		r6 := &sendBtc{
			Address: "ch98fh6y13905yhg190h6198yhtgnb",
			Amount:  "100",
		}
		err := r6.Validate()
		assert.NoError(t, err)
	})
}
