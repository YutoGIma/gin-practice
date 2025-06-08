package main

import (
	"myapp/app/config"
	"myapp/app/controller"
	"myapp/app/infra"
	"myapp/app/middleware"
	"myapp/app/routes"
	"myapp/app/service"
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

	// Services
	baseService := service.NewBaseService(db)

	// Use cases
	baseUseCase := usecase.NewBaseUseCase(baseService)

	// Controllers
	baseController := controller.NewBaseController(baseUseCase)

	// Router setup
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger(logger))
	r.Use(middleware.ErrorHandler())

	r = routes.SetupRouter(baseController)

	r.Run(":" + cfg.ServerPort)
}
