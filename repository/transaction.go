package repository

// You only need **one** of these per package!
// see https://github.com/maxbrunsfeld/counterfeiter#step-2b---add-counterfeitergenerate-directives
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"context"
	"time"

	"anylogibtc/dto"
)

// Transaction is the repository interface to fulfill to use the wallet aggregate
//
//counterfeiter:generate . TransactionRepository
type TransactionRepository interface {
	Send(ctx context.Context, transaction dto.TransactionDTO) error
	History(ctx context.Context, from time.Time, to time.Time) ([]dto.TransactionDTO, error)
}
