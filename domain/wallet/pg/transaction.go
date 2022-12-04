package pg

import (
	"anylogibtc/entity"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgTransaction struct {
	dbpool *pgxpool.Pool
}

func NewPgTransaction(ctx context.Context, db *pgxpool.Pool) *PgTransaction {
	return &PgTransaction{
		dbpool: db,
	}
}

func (pt *PgTransaction) Send(ctx context.Context, t entity.Transaction) error {
	return nil
}

func (pt *PgTransaction) History(ctx context.Context, id int) ([]entity.Transaction, error) {
	return nil, nil
}
