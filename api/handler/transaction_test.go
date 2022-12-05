package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"anylogibtc/api/handler"
	"anylogibtc/dto"
	"anylogibtc/services/transaction"
	"anylogibtc/services/transaction/transactionfakes"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestSaveTransaction(t *testing.T) {
	saveJSON := `{"datetime": "2019-10-05T14:45:05+07:00","amount": 10}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/wallets", strings.NewReader(saveJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	trxService := &transactionfakes.FakeTransactionService{}
	trxHandler := handler.NewTransactionHandler(trxService)
	// Assertions
	if assert.NoError(t, trxHandler.Save(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.JSONEq(t, `{"message":"data created successfully"}`, rec.Body.String())
	}
}

func TestSaveFailTransaction(t *testing.T) {
	saveJSON := `{"datetime": "2019-10-05T14:45:05+07:00","amount": 10}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/wallets", strings.NewReader(saveJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	trxService := &transactionfakes.FakeTransactionService{}
	trxService.SendReturns(errors.New("ups error"))
	trxHandler := handler.NewTransactionHandler(trxService)
	// Assertions
	if assert.Error(t, trxHandler.Save(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"error":"ups error"}`, rec.Body.String())
	}

	req = httptest.NewRequest(http.MethodPost, "/v1/wallets", strings.NewReader(`{"datetime": "2019-10-05T14:45:05+07:00","amount": 0}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	// Assertions
	if assert.Error(t, trxHandler.Save(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"error":"amount can't be 0"}`, rec.Body.String())
	}

	reqInvalidObject := httptest.NewRequest(http.MethodPost, "/v1/wallets", strings.NewReader(`{"datetime": "2019-10-05T14:45:05+07:00"`))
	reqInvalidObject.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recInvalidObject := httptest.NewRecorder()
	c1 := e.NewContext(reqInvalidObject, recInvalidObject)
	// Assertions
	if assert.Error(t, trxHandler.Save(c1)) {
		assert.Equal(t, http.StatusBadRequest, recInvalidObject.Code)
	}
}

func TestGetHistory(t *testing.T) {
	paramsHistory := `{"startDatetime":"2011-10-05T10:48:01+00:00","endDatetime":"2011-10-05T18:48:02+00:00"}`
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/wallets/history", strings.NewReader(paramsHistory))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	trxService := &transactionfakes.FakeTransactionService{}
	histories := transaction.HistoriesDTO{}
	histories = append(histories, dto.TransactionDTO{Amount: decimal.NewFromFloat(1.2), Datetime: time.Now()})
	rawJson, _ := json.Marshal(histories)
	trxService.HistoryReturns(histories, nil)
	trxHandler := handler.NewTransactionHandler(trxService)
	// Assertions
	if assert.NoError(t, trxHandler.History(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(rawJson), rec.Body.String())
	}

	req = httptest.NewRequest(http.MethodPost, "/v1/wallets/history", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	// Assertions
	if assert.NoError(t, trxHandler.History(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, string(rawJson), rec.Body.String())
	}

	req = httptest.NewRequest(http.MethodPost, "/v1/wallets/history", strings.NewReader("as{"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	// Assertions
	if assert.Error(t, trxHandler.History(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	req = httptest.NewRequest(http.MethodPost, "/v1/wallets/history", strings.NewReader(paramsHistory))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	trxService.HistoryReturns(transaction.HistoriesDTO{}, errors.New("ups error"))
	// Assertions
	if assert.Error(t, trxHandler.History(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
