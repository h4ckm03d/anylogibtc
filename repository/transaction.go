package repository

import "anylogibtc/ent"

// Transaction is the repository interface to fulfill to use the wallet aggregate
//
//counterfeiter:generate . Transaction
type Transaction interface {
	Send(transaction ent.Transaction) error
	History(id int) ([]ent.Transaction, error)
}
