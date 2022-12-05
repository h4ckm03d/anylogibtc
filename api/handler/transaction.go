package handler

import (
	"errors"
	"net/http"
	"time"

	"anylogibtc/dto"
	"anylogibtc/services/transaction"

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
		c.JSON(http.StatusBadRequest, dto.NewResponse("", err.Error()))
		return err
	}

	if input.Amount.Equal(decimal.NewFromInt(0)) {
		err := errors.New("amount can't be 0")
		c.JSON(http.StatusBadRequest, dto.NewResponse("", err.Error()))
		return err
	}

	if err := t.serviceTransaction.Send(c.Request().Context(), *input); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewResponse("", err.Error()))
		return err
	}

	return c.JSON(http.StatusCreated, dto.NewResponse("data created successfully", ""))
}

func (t *TransactionHandler) History(ctx echo.Context) error {
	params := new(dto.HistoryParamsDTO)
	if err := ctx.Bind(params); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewResponse("", err.Error()))
		return err
	}

	if params.StartDatetime.IsZero() {
		params.StartDatetime = time.Now().Add(time.Hour * -1)
	}

	if params.EndDatetime.IsZero() {
		params.EndDatetime = time.Now()
	}

	res, err := t.serviceTransaction.History(ctx.Request().Context(), params.StartDatetime, params.EndDatetime)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewResponse("", err.Error()))
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}
