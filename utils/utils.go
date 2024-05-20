package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"strconv"

	"github.com/ownxgv/Genesis_SES_task/configs"
	"github.com/ownxgv/Genesis_SES_task/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.AppConfig.DBUser,
		configs.AppConfig.DBPassword,
		configs.AppConfig.DBHost,
		configs.AppConfig.DBPort,
		configs.AppConfig.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.CurrencyRate{}, &models.Subscription{})
}

type CoinbaseResponse struct {
	Data struct {
		Currency string `json:"currency"`
		Rates    struct {
			USD string `json:"USD"`
		} `json:"rates"`
	} `json:"data"`
}

func FetchUSDRate() (float64, error) {
	resp, err := http.Get("https://api.coinbase.com/v2/exchange-rates?currency=UAH")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch exchange rates, status: %d", resp.StatusCode)
	}

	var coinbaseResponse CoinbaseResponse
	if err := json.NewDecoder(resp.Body).Decode(&coinbaseResponse); err != nil {
		return 0, err
	}

	rate, err := strconv.ParseFloat(coinbaseResponse.Data.Rates.USD, 64)
	if err != nil {
		return 0, err
	}

	return rate, nil
}

func SendEmail(to string, rate float64) error {
	auth := smtp.PlainAuth("", configs.AppConfig.EmailSender, configs.AppConfig.EmailPass, "smtp.gmail.com")

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: Daily Currency Rate Update\r\n\r\nCurrent UAH to USD rate: %.2f\r\n", configs.AppConfig.EmailSender, to, rate)

	err := smtp.SendMail("smtp.gmail.com:587", auth, configs.AppConfig.EmailSender, []string{to}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}
