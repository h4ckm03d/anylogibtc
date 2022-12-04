package pg

import (
	"anylogibtc/dto"
	"anylogibtc/entity"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgTransaction struct {
	dbpool *pgxpool.Pool
}

func NewPgTransaction(db *pgxpool.Pool) *PgTransaction {
	return &PgTransaction{
		dbpool: db,
	}
}

func (pt *PgTransaction) Send(ctx context.Context, t dto.TransactionDTO) error {
	if _, err := pt.dbpool.Exec(ctx, "insert into transactions(datetime, amount) values ($1, $2)", t.Datetime, t.Amount); err != nil {
		return err
	}
	return nil
}

func (pt *PgTransaction) History(ctx context.Context, from time.Time, to time.Time) ([]entity.Transaction, error) {
	return nil, nil
}
