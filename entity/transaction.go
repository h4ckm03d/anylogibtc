package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Datetime holds the value of the "datetime" field.
	Datetime time.Time `json:"datetime,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
}
