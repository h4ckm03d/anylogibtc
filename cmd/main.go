package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"anylogibtc/api/handler"

	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	DatabaseURL string
	Port        int
}

func DefaultConfig() *Config {
	port, err := strconv.Atoi(Default(os.Getenv("PORT"), "3000"))
	if err != nil {
		log.Fatalf("invalid port: %v", err)
	}
	return &Config{
		DatabaseURL: Default(os.Getenv("DATABASE_URL"), "postgres://timescaledb:secret@localhost:5432/anylogi?sslmode=disable"),
		Port:        port,
	}
}

func Default(input string, defaultValue string) string {
	if input == "" {
		return defaultValue
	}

	return input
}

type Command struct {
	server handler.Server
	db     *pgxpool.Pool
}

func NewCommand(cfg *Config) *Command {
	poolConfig, err := pgxpool.ParseConfig(cfg.DatabaseURL)
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
		server: handler.NewEchoServer(cfg.Port, db),
		db:     db,
	}
}

func (cmd *Command) Run() {
	cmd.server.Run()
}

func main() {
	config := DefaultConfig()
	fmt.Println(config)
	cmd := NewCommand(config)
	cmd.Run()
}
