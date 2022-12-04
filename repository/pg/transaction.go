package pg

import (
	"anylogibtc/dto"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
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

func (pt *PgTransaction) History(ctx context.Context, from time.Time, to time.Time) ([]dto.TransactionDTO, error) {
	results := []dto.TransactionDTO{}
	query := `WITH
	history as (
	  select
		hour as dt,
		sum(total) over (
		  order by
			hour asc rows between unbounded preceding
			and current row
		) as total
	  from
		transaction_hourly
	)
  select
	dt, total
  from
	history
  WHERE
  	dt >= $1
	AND dt < $2`

	rows, err := pt.dbpool.Query(ctx, query, from, to)

	if err != nil {
		return results, err
	}

	for rows.Next() {
		var amount decimal.Decimal
		var datetime time.Time
		err = rows.Scan(&datetime, &amount)
		if err != nil {
			return results, err
		}
		results = append(results, dto.TransactionDTO{Datetime: datetime.UTC(), Amount: amount})
	}

	return results, rows.Err()
}
