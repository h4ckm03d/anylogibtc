package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"anylogibtc/api/handler"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthz(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(t, handler.Healthz(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "ok", rec.Body.String())
	}
}
