package pg

import (
	"anylogibtc/ent"
	"context"
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
)

type PgTransaction struct {
	client *ent.Client
}

func NewPgTransaction(ctx context.Context, db *sql.DB) *PgTransaction {
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)

	return &PgTransaction{
		client: ent.NewClient(ent.Driver(drv)),
	}
}

func (pt *PgTransaction) Send(ctx context.Context, t ent.Transaction) error {
	return nil
}

func (pt *PgTransaction) History(ctx context.Context, id int) ([]ent.Transaction, error) {
	return nil, nil
}
