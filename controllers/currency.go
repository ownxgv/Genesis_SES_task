package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/your-username/currency-service/models"
	"github.com/your-username/currency-service/services"
	"net/http"
)

type CurrencyController struct {
	Service services.CurrencyService
}

func NewCurrencyController(service services.CurrencyService) CurrencyController {
	return CurrencyController{Service: service}
}

// GetCurrencyRate godoc
// @Summary Get current UAH to USD rate
// @Description Get the current exchange rate of UAH to USD
// @Tags currency
// @Produce json
// @Success 200 {object} models.CurrencyRate
// @Router /currency [get]
func (c *CurrencyController) GetCurrencyRate(ctx *gin.Context) {
	rate, err := c.Service.GetCurrencyRate()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, rate)
}

// SubscribeEmail godoc
// @Summary Subscribe to currency rate updates
// @Description Subscribe to receive daily emails with the updated UAH to USD rate
// @Tags subscription
// @Param email body string true "Email address"
// @Success 200 {string} string "ok"
// @Failure 400 {object} models.ErrorResponse
// @Router /subscribe [post]
func (c *CurrencyController) SubscribeEmail(ctx *gin.Context) {
	var req models.SubscriptionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	err := c.Service.SubscribeEmail(req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "ok")
}
