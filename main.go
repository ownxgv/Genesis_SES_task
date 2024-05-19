package main

import (
	"github.com/robfig/cron"
	"github.com/your-username/currency-service/routes"
	"github.com/your-username/currency-service/utils"
	"log"
)

// @title Currency Service API
// @version 1.0
// @description API for getting currency rates and subscribing to email notifications.

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// Load configurations
	configs.LoadConfig()

	// Connect to the database
	db := utils.ConnectDB()

	// Run database migrations
	utils.MigrateDB(db)

	// Create a new router
	router := routes.NewRouter(db)

	// Serve the API
	log.Printf("Server started on port %s", configs.Config.Port)
	err := router.Run(":" + configs.Config.Port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	c := cron.New()
	_, err = c.AddFunc("0 8 * * *", func() {
		err := currencyService.SendDailyRates()
		if err != nil {
			log.Printf("Failed to send daily rates: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("Failed to create cron job: %v", err)
	}

	c.Start()

	// Запуск HTTP-сервера
	// ...
}
