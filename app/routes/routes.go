package routes

import (
	"myapp/app/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(baseController *controller.BaseController) *gin.Engine {
	r := gin.Default()
	r.GET("/users", baseController.UserController.GetUsers)
	r.GET("/users/:id", baseController.UserController.GetUserDetail)
	r.POST("/users", baseController.UserController.CreateUser)
	r.PUT("/users/:id", baseController.UserController.UpdateUser)
	r.DELETE("/users/:id", baseController.UserController.DeleteUser)
	r.GET("/products", baseController.ProductController.GetProducts)
	r.POST("/products", baseController.ProductController.CreateProduct)
	r.PUT("/products/:id", baseController.ProductController.UpdateProduct)
	r.DELETE("/products/:id", baseController.ProductController.DeleteProduct)
	r.GET("/inventories", baseController.InventoryController.GetInventories)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
