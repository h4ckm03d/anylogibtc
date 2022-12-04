package main

import (
	"log"
	"os"
	"strconv"
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

func main() {
	config := DefaultConfig()

	cmd := NewCommand(config)
	cmd.Run()
}
