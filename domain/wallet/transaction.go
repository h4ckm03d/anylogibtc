package wallet

import (
	"anylogibtc/dto"
	"anylogibtc/entity"
	"context"
	"time"
)

// Transaction is the repository interface to fulfill to use the wallet aggregate
//
//counterfeiter:generate . TransactionRepository
type TransactionRepository interface {
	Send(ctx context.Context, transaction dto.TransactionDTO) error
	History(ctx context.Context, from time.Time, to time.Time) ([]entity.Transaction, error)
}
