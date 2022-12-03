package wallet

import (
	"errors"

	"anylogibtc/ent"
)

var (
	// ErrWalletNotFound is returned when a wallet is not found
	ErrWalletNotFound = errors.New("the wallet was not found")
	// ErrWalletAlreadyExist is returned when trying to add a wallet that already exists
	ErrWalletAlreadyExist = errors.New("the wallet already exists")
)

// Wallet is the repository interface to fulfill to use the wallet aggregate
//
//counterfeiter:generate . WalletRepository
type WalletRepository interface {
	GetAll() ([]ent.Wallet, error)
	GetByID(id int) (ent.Wallet, error)
	Add(wallet ent.Wallet) error
	Update(wallet ent.Wallet) error
	Delete(id int) error
}
