// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ownxgv/Genesis_SES_task/controllers"
	"github.com/ownxgv/Genesis_SES_task/docs"
	"github.com/ownxgv/Genesis_SES_task/repositories"
	"github.com/ownxgv/Genesis_SES_task/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// NewRouter создает новый маршрутизатор Gin с настроенными маршрутами API.
func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Инициализация репозиториев
	currencyRepo := repositories.NewCurrencyRepository(db)

	// Инициализация сервисов
	currencyService := services.NewCurrencyService(*currencyRepo)

	// Инициализация контроллеров
	currencyController := controllers.NewCurrencyController(currencyService)

	// Документация Swagger
	docs.SwaggerInfo.Host = "localhost:8080"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Маршруты API
	v1 := router.Group("/api/v1")
	{
		v1.GET("/currency", currencyController.GetCurrencyRate)
		v1.POST("/subscribe", currencyController.SubscribeEmail)
	}

	return router
}
