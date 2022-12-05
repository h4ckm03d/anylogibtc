//go:build integration

package pg_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"anylogibtc/dto"
	"anylogibtc/repository/pg"
	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

func TestPgTransaction(t *testing.T) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://timescaledb:secret@localhost:5432/anylogi_test?sslmode=disable"
	}
	poolConfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatalln("Unable to parse DATABASE_URL:", err)
	}
	poolConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxdecimal.Register(conn.TypeMap())
		return nil
	}

	poolConfig.ConnConfig.RuntimeParams["timezone"] = "UTC"

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		t.Fatalf("Unable to create connection pool: %v", err)
	}

	// truncate database testing
	if _, err := db.Query(context.Background(), "TRUNCATE TABLE transactions"); err != nil {
		t.Fatalf("Unable to truncate table: %v", err)
	}

	type fields struct {
		dbpool *pgxpool.Pool
	}
	type args struct {
		action func(ctx context.Context, pt *pg.PgTransaction, t *testing.T) error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test Send 1.2",
			fields: fields{
				dbpool: db,
			},
			args: args{
				action: func(ctx context.Context, pt *pg.PgTransaction, t *testing.T) error {
					data := dto.TransactionDTO{
						Amount:   decimal.NewFromFloat(1.1),
						Datetime: time.Now().Add(time.Hour * -2).UTC(),
					}
					return pt.Send(ctx, data)
				},
			},
			wantErr: false,
		},
		{
			name: "Test Send 1.1",
			fields: fields{
				dbpool: db,
			},
			args: args{
				action: func(ctx context.Context, pt *pg.PgTransaction, t *testing.T) error {
					data := dto.TransactionDTO{
						Amount:   decimal.NewFromFloat(1.2),
						Datetime: time.Now().UTC(),
					}
					return pt.Send(ctx, data)
				},
			},
			wantErr: false,
		},
		{
			name: "Test Send 0.9",
			fields: fields{
				dbpool: db,
			},
			args: args{
				action: func(ctx context.Context, pt *pg.PgTransaction, t *testing.T) error {
					data := dto.TransactionDTO{
						Amount:   decimal.NewFromFloat(1.2),
						Datetime: time.Now().UTC(),
					}
					return pt.Send(ctx, data)
				},
			},
			wantErr: false,
		},
		{
			name: "get history",
			fields: fields{
				dbpool: db,
			},
			args: args{
				action: func(ctx context.Context, pt *pg.PgTransaction, t *testing.T) error {
					now := time.Now().UTC()
					from := now.Add(time.Hour * -3)
					results, err := pt.History(ctx, from, now)
					if len(results) != 2 {
						t.Errorf("Expected 2 results, got %d. %v", len(results), results)
					}
					t.Logf("got data : %v", results)
					return err
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := pg.NewPgTransaction(tt.fields.dbpool)
			if err := tt.args.action(context.Background(), pt, t); (err != nil) != tt.wantErr {
				t.Errorf("PgTransaction.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
