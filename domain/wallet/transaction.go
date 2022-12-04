package wallet

import (
	"anylogibtc/entity"
	"context"
)

// Transaction is the repository interface to fulfill to use the wallet aggregate
//
//counterfeiter:generate . TransactionRepository
type TransactionRepository interface {
	Send(ctx context.Context, transaction entity.Transaction) error
	History(ctx context.Context, id int) ([]entity.Transaction, error)
}
