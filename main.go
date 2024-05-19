package main

import (
	"log"
	"time"

	"github.com/ownxgv/Genesis_SES_task/configs"
	"github.com/ownxgv/Genesis_SES_task/repositories"
	"github.com/ownxgv/Genesis_SES_task/routes"
	"github.com/ownxgv/Genesis_SES_task/services"
	"github.com/ownxgv/Genesis_SES_task/utils"
)

func main() {
	// Load configurations
	if err := configs.LoadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to the database
	db := utils.ConnectDB()

	// Run database migrations
	utils.MigrateDB(db)

	// Создание канала для ошибок
	errChan := make(chan error)

	// Создание сервиса валют
	currencyRepo := repositories.NewCurrencyRepository(db)
	var currencyService services.CurrencyService = services.NewCurrencyService(*currencyRepo)

	// Запуск горутины для отправки ежедневных уведомлений
	go func() {
		ticker := time.NewTicker(24 * time.Hour) // каждые 24 часа
		for {
			select {
			case <-ticker.C:
				if err := currencyService.SendDailyRates(); err != nil {
					errChan <- err
				}
			}
		}
	}()

	// Запуск обработчика ошибок в отдельной горутине
	go func() {
		for err := range errChan {
			if err != nil {
				log.Printf("Failed to send daily rates: %v", err)
			}
		}
	}()

	// Create a new router
	router := routes.NewRouter(db)

	// Serve the API
	log.Printf("Server started on port %s", configs.AppConfig.Port)
	if err := router.Run(":" + configs.AppConfig.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
