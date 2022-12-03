package wallet

import (
	"anylogibtc/ent"
	"context"
)

// Transaction is the repository interface to fulfill to use the wallet aggregate
//
//counterfeiter:generate . TransactionRepository
type TransactionRepository interface {
	Send(ctx context.Context, transaction ent.Transaction) error
	History(ctx context.Context, id int) ([]ent.Transaction, error)
}
