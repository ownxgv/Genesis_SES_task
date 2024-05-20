// controllers/currency_controller_test.go

package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ownxgv/Genesis_SES_task/models"
	mocks "github.com/ownxgv/Genesis_SES_task/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrencyRate(t *testing.T) {
	mockService := &mocks.MockCurrencyService{}
	mockService.On("GetCurrencyRate").Return(&models.CurrencyRate{USDRate: 27.5}, nil)

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
	mockService := &mocks.MockCurrencyService{}
	mockService.On("SubscribeEmail", "test@example.com").Return(nil)

	controller := NewCurrencyController(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	subscription := models.Subscription{Email: "test@example.com"}
	jsonValue, _ := json.Marshal(subscription)
	c.Request, _ = http.NewRequest("POST", "/subscribe", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	controller.SubscribeEmail(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Email subscribed successfully", response["message"])
}
