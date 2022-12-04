package handler

import (
	"anylogibtc/domain/wallet/pg"
	"anylogibtc/services/transaction"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Server interface {
	Run()
}

type EchoServer struct {
	e    *echo.Echo
	Port int
	db   *pgxpool.Pool
}

func NewEchoServer(port int, db *pgxpool.Pool) *EchoServer {
	return &EchoServer{
		e:    echo.New(),
		Port: port,
		db:   db,
	}
}

func (s *EchoServer) SetupRoutes() {

	transactionRepo := pg.NewPgTransaction(s.db)
	transactionService := transaction.NewTransactionService(transactionRepo)
	transactionHandler := NewTransactionHandler(transactionService)

	g := s.e.Group("/v1")
	// health check routes
	g.GET("/healthz", Healthz)
	g.POST("/wallets", transactionHandler.Save)
	g.GET("/wallets", transactionHandler.History)
}

func (s *EchoServer) Run() {
	s.e.Logger.SetLevel(log.INFO)

	s.SetupRoutes()

	// Start server
	go func() {
		if err := s.e.Start(fmt.Sprintf(":%d", s.Port)); err != nil && err != http.ErrServerClosed {
			s.e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.e.Shutdown(ctx); err != nil {
		s.e.Logger.Fatal(err)
	}
}

func Healthz(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
