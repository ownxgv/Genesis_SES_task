// controllers/currency_test.go
package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ownxgv/Genesis_SES_task/mocks"
	"github.com/ownxgv/Genesis_SES_task/models"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrencyRate(t *testing.T) {
	mockService := mocks.NewMockCurrencyService()
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
	mockService := mocks.NewMockCurrencyService()
	controller := NewCurrencyController(mockService)

	// Создаем тестовый запрос
	reqBody := map[string]string{"email": "test@example.com"}
	reqBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBytes))
	req.Header.Set("Content-Type", "application/json")

	// Запускаем контроллер
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	controller.SubscribeEmail(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Email subscribed successfully", response["message"])
}
