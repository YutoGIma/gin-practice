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

// @title           Gin Practice API
// @version         1.0
// @description     商品在庫管理システムのRESTful API
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @schemes http
func main() {
	cfg := config.NewConfig()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	db := infra.SetupDB()
	infra.DBMigration(db)

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
