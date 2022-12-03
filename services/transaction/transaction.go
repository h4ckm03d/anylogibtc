package transaction

import (
	"time"

	"github.com/shopspring/decimal"
)

type TransactionDTO struct {
	// Datetime holds the value of the "datetime" field.
	Datetime time.Time `json:"datetime,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
	// SenderID holds the value of the "sender_id" field.
	SenderID int `json:"sender_id,omitempty"`
	// RecipientID holds the value of the "recipient_id" field.
	RecipientID int `json:"recipient_id,omitempty"`
}

type HistoryDTO struct {
	// Datetime holds the value of the "datetime" field.
	Datetime time.Time `json:"datetime,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
}

type HistoryParams struct {
	WalletId int `json:"walletId,omitempty"`
	From     time.Time
	To       time.Time
}

type HistoriesDTO []HistoryDTO

type Transaction interface {
	Send(td TransactionDTO) error
	History(p HistoryParams) HistoriesDTO
}
