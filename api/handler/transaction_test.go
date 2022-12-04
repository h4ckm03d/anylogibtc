package handler_test

import (
	"anylogibtc/api/handler"
	"anylogibtc/services/transaction/transactionfakes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSaveTransaction(t *testing.T) {
	var (
		saveJSON = `{"datetime": "2019-10-05T14:45:05+07:00","amount": 10}`
	)

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/wallets", strings.NewReader(saveJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	trxService := &transactionfakes.FakeTransaction{}
	trxHandler := handler.NewTransactionHandler(trxService)
	// Assertions
	if assert.NoError(t, trxHandler.Save(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.JSONEq(t, `{"message":"data created successfully"}`, rec.Body.String())
	}
}

func TestSaveFailTransaction(t *testing.T) {
	var (
		saveJSON = `{"datetime": "2019-10-05T14:45:05+07:00","amount": 10}`
	)

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/wallets", strings.NewReader(saveJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	trxService := &transactionfakes.FakeTransaction{}
	trxService.SendReturns(errors.New("ups error"))
	trxHandler := handler.NewTransactionHandler(trxService)
	// Assertions
	if assert.NoError(t, trxHandler.Save(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"error":"ups error"}`, rec.Body.String())
	}

	req = httptest.NewRequest(http.MethodPost, "/v1/wallets", strings.NewReader(`{"datetime": "2019-10-05T14:45:05+07:00","amount": 0}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	// Assertions
	if assert.NoError(t, trxHandler.Save(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"error":"amount can't be 0"}`, rec.Body.String())
	}

	reqInvalidObject := httptest.NewRequest(http.MethodPost, "/v1/wallets", strings.NewReader(`{"datetime": "2019-10-05T14:45:05+07:00"`))
	reqInvalidObject.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recInvalidObject := httptest.NewRecorder()
	c1 := e.NewContext(reqInvalidObject, recInvalidObject)
	// Assertions
	if assert.NoError(t, trxHandler.Save(c1)) {
		assert.Equal(t, http.StatusBadRequest, recInvalidObject.Code)
	}

}
