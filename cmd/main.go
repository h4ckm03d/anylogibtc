package main

import (
	"anylogibtc/api/handler"
	"anylogibtc/ent"
	"context"
	"database/sql"
	"flag"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Config struct {
	DatabaseURL   string
	WithMigration bool
	Port          int
}

func main() {
	config := Config{
		DatabaseURL:   "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable",
		WithMigration: false,
		Port:          3000,
	}
	flag.IntVar(&config.Port, "p", 3000, "number of lines to read from the file")
	flag.StringVar(&config.DatabaseURL, "c", "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable", "database connection string")
	flag.BoolVar(&config.WithMigration, "migrate", false, "run migration")
	flag.Parse()

	cmd := NewCommand(config)
	if config.WithMigration {
		log.Println("run migration")
		if err := cmd.Migrate(); err != nil {
			log.Fatal(err)
		}
	}
	cmd.Run()
}

type Command struct {
	server    handler.Server
	entClient *ent.Client
}

func NewCommand(cfg Config) *Command {
	return &Command{
		server:    handler.NewEchoServer(cfg.Port),
		entClient: Open(cfg.DatabaseURL),
	}
}

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func (cmd *Command) Migrate() error {
	ctx := context.Background()
	if err := cmd.entClient.Schema.Create(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *Command) Run() {
	cmd.server.Run()
}

type Commander interface {
	Migrate() error
	Run()
}
