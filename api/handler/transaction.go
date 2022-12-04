package handler

import (
	"anylogibtc/dto"
	"anylogibtc/services/transaction"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

type TransactionHandler struct {
	serviceTransaction transaction.TransactionService
}

func NewTransactionHandler(serviceTransaction transaction.TransactionService) *TransactionHandler {

	return &TransactionHandler{
		serviceTransaction: serviceTransaction,
	}
}

func (t *TransactionHandler) Save(c echo.Context) error {
	input := new(dto.TransactionDTO)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, dto.NewResponse("", err.Error()))
	}

	if input.Amount.Equal(decimal.NewFromInt(0)) {
		return c.JSON(http.StatusBadRequest, dto.NewResponse("", "amount can't be 0"))
	}

	if err := t.serviceTransaction.Send(c.Request().Context(), *input); err != nil {
		return c.JSON(http.StatusBadRequest, dto.NewResponse("", err.Error()))
	}

	return c.JSON(http.StatusCreated, dto.NewResponse("data created successfully", ""))
}

func (t *TransactionHandler) History(ctx echo.Context) error {
	return nil
}
