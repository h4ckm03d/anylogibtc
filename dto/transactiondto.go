package dto

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

type HistoryParamsDTO struct {
	StartDatetime time.Time `json:"startDatetime"`
	EndDatetime   time.Time `json:"endDatetime"`
}

type HistoryDTO struct {
	// Datetime holds the value of the "datetime" field.
	Datetime time.Time `json:"datetime"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount"`
}

type ResponseDTO struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewResponse(message string, error string) ResponseDTO {
	return ResponseDTO{
		Message: message,
		Error:   error,
	}
}
