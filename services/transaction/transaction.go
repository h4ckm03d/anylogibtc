package transaction

// You only need **one** of these per package!
// see https://github.com/maxbrunsfeld/counterfeiter#step-2b---add-counterfeitergenerate-directives
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"anylogibtc/domain/wallet"
	"anylogibtc/dto"
	"context"
	"time"
)

type HistoriesDTO []dto.HistoryDTO

//counterfeiter:generate . TransactionService
type TransactionService interface {
	Send(ctx context.Context, td dto.TransactionDTO) error
	History(ctx context.Context, from time.Time, to time.Time) (HistoriesDTO, error)
}

type transaction struct {
	repo wallet.TransactionRepository
}

func NewTransactionService(repo wallet.TransactionRepository) TransactionService {
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
	histories, err := t.repo.History(ctx, from, to)
	results := HistoriesDTO{}
	if err != nil {
		return results, err
	}

	for _, history := range histories {
		results = append(results, dto.HistoryDTO{
			Datetime: history.Datetime,
			Amount:   history.Amount,
		})
	}

	return results, nil
}
