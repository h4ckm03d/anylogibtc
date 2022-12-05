package transaction

// You only need **one** of these per package!
// see https://github.com/maxbrunsfeld/counterfeiter#step-2b---add-counterfeitergenerate-directives
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"context"
	"errors"
	"time"

	"anylogibtc/dto"
	"anylogibtc/repository"
)

type HistoriesDTO []dto.TransactionDTO

//counterfeiter:generate . TransactionService
type TransactionService interface {
	Send(ctx context.Context, td dto.TransactionDTO) error
	History(ctx context.Context, from time.Time, to time.Time) (HistoriesDTO, error)
}

type transaction struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transaction{repo: repo}
}

func (t *transaction) Send(ctx context.Context, td dto.TransactionDTO) error {
	err := t.repo.Send(ctx, td)
	if err != nil {
		return err
	}
	return nil
}

func (t *transaction) History(ctx context.Context, from time.Time, to time.Time) (HistoriesDTO, error) {
	results := HistoriesDTO{}

	if from.UTC().After(to.UTC()) {
		return results, errors.New("from datetime cannot be after to datetime")
	}

	histories, err := t.repo.History(ctx, from.UTC(), to.UTC())
	if err != nil {
		return results, err
	}

	return HistoriesDTO(histories), nil
}
