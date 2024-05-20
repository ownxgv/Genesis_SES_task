package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ownxgv/Genesis_SES_task/models"
	"github.com/ownxgv/Genesis_SES_task/services"
)

type CurrencyController struct {
	Service services.CurrencyService
}

func NewCurrencyController(service services.CurrencyService) *CurrencyController {
	return &CurrencyController{Service: service}
}

func (cc *CurrencyController) GetCurrencyRate(c *gin.Context) {
	rate, err := cc.Service.GetCurrencyRate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rate)
}

func (cc *CurrencyController) SubscribeEmail(c *gin.Context) {
	var subscription models.Subscription
	if err := c.BindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cc.Service.SubscribeEmail(subscription.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email subscribed successfully"})
}
