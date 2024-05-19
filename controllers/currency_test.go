// controllers/currency_test.go
package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/your-username/currency-service/models"
	"github.com/your-username/currency-service/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCurrencyRate(t *testing.T) {
	mockService := services.NewMockCurrencyService()
	controller := NewCurrencyController(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	controller.GetCurrencyRate(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var rate models.CurrencyRate
	err := json.Unmarshal(w.Body.Bytes(), &rate)
	assert.NoError(t, err)
	assert.Equal(t, 27.5, rate.USDRate)
}

func TestSubscribeEmail(t *testing.T) {
	mockService := services.NewMockCurrencyService()
	controller := NewCurrencyController(mock