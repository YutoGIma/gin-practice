package main

import (
	"myapp/app/config"
	"myapp/app/controller"
	"myapp/app/infra"
	"myapp/app/middleware"
	"myapp/app/routes"
	"myapp/app/usecase"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg := config.NewConfig()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	db := infra.SetupDB()
	db.AutoMigrate()

	// Use cases
	productUseCase := usecase.NewProductUseCase(db)
	tenantUseCase := usecase.NewTenantUseCase(db)
	inventoryUseCase := usecase.NewInventoryUseCase(db)
	userUseCase := usecase.NewUserUseCase(db)

	// Controllers
	productController := controller.NewProductController(productUseCase)

	// Router setup
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger(logger))
	r.Use(middleware.ErrorHandler())

	routes.SetupRouter(r, productController)

	r.Run(":" + cfg.ServerPort)
}
