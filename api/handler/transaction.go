package handler

import (
	"anylogibtc/services/transaction"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

type TransactionHandler struct {
	serviceTransaction transaction.Transaction
}

func NewTransactionHandler(serviceTransaction transaction.Transaction) *TransactionHandler {

	return &TransactionHandler{
		serviceTransaction: serviceTransaction,
	}
}

func (t *TransactionHandler) Save(c echo.Context) error {
	input := new(transaction.TransactionDTO)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseDTO{Error: err.Error()})
	}

	if input.Amount.Equal(decimal.NewFromInt(0)) {
		return c.JSON(http.StatusBadRequest, ResponseDTO{Error: "amount can't be 0"})
	}

	if err := t.serviceTransaction.Send(*input); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseDTO{Error: err.Error()})
	}

	return c.JSON(http.StatusCreated, ResponseDTO{Message: "data created successfully"})
}

func (t *TransactionHandler) History(ctx echo.Context) error {
	return nil
}
