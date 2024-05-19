// utils/utils.go
package utils

import (
	"encoding/json"
	"fmt"
	"github.com/your-username/currency-service/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"net/smtp"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.Config.DBUser,
		configs.Config.DBPassword,
		configs.Config.DBHost,
		configs.Config.DBPort,
		configs.Config.DBName,
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

type ExchangeRate struct {
	CurrencyCode string  `json:"cc"`
	Rate         float64 `json:"rate"`
}

func FetchUSDRate() (float64, error) {
	// Запрос к API НБУ для получения курсов валют
	resp, err := http.Get("https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?json")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch exchange rates, status: %d", resp.StatusCode)
	}

	// Декодируем JSON-ответ
	var exchangeRates []ExchangeRate
	if err := json.NewDecoder(resp.Body).Decode(&exchangeRates); err != nil {
		return 0, err
	}

	// Находим курс гривны к доллару
	for _, rate := range exchangeRates {
		if rate.CurrencyCode == "USD" {
			return rate.Rate, nil
		}
	}
	return 0, fmt.Errorf("USD rate not found in the response")
}

func SendEmail(to string, rate float64) error {
	auth := smtp.PlainAuth("", configs.Config.EmailSender, configs.Config.EmailPass, "smtp.gmail.com")

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: Daily Currency Rate Update\r\n\r\nCurrent UAH to USD rate: %.2f\r\n", configs.Config.EmailSender, to, rate)

	err := smtp.SendMail("smtp.gmail.com:587", auth, configs.Config.EmailSender, []string{to}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}
