// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/your-username/currency-service/controllers"
	"github.com/your-username/currency-service/docs"
	"github.com/your-username/currency-service/repositories"
	"github.com/your-username/currency-service/services"
	"gorm.io/gorm"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Initialize repositories
	currencyRepo := repositories.NewCurrencyRepository(db)

	// Initialize services
	currencyService := services.NewCurrencyService(*currencyRepo)

	// Initialize controllers
	currencyController := controllers.NewCurrencyController(currencyService)

	// Swagger documentation
	docs.SwaggerInfo.Host = "localhost:8080"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/currency", currencyController.GetCurrencyRate)
		v1.POST("/subscribe", currencyController.SubscribeEmail)
	}

	return router
}
