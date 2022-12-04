package main

import (
	"anylogibtc/api/handler"
	"context"
	"log"
	"os"

	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Command struct {
	server handler.Server
	db     *pgxpool.Pool
}

func NewCommand(cfg *Config) *Command {
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
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
		log.Fatalln("Unable to create connection pool:", err)
	}

	return &Command{
		server: handler.NewEchoServer(cfg.Port),
		db:     db,
	}
}

func (cmd *Command) Run() {
	cmd.server.Run()
}
