package transaction

// You only need **one** of these per package!
// see https://github.com/maxbrunsfeld/counterfeiter#step-2b---add-counterfeitergenerate-directives
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"time"

	"github.com/shopspring/decimal"
)

type TransactionDTO struct {
	// Datetime holds the value of the "datetime" field.
	Datetime time.Time `json:"datetime" `
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount" `
}

type HistoryDTO struct {
	// Datetime holds the value of the "datetime" field.
	Datetime time.Time `json:"datetime"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount"`
}

type HistoryParams struct {
	WalletId int `json:"walletId,omitempty"`
	From     time.Time
	To       time.Time
}

type HistoriesDTO []HistoryDTO

//counterfeiter:generate . Transaction
type Transaction interface {
	Send(td TransactionDTO) error
	History(p HistoryParams) HistoriesDTO
}
