package routes

import (
	"myapp/app/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	
	_ "myapp/docs" // swagger docs
)

func SetupRouter(baseController controller.BaseController) *gin.Engine {
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
	r.POST("/inventories", baseController.InventoryController.CreateInventory)
	r.PUT("/inventories/:id", baseController.InventoryController.UpdateInventory)
	r.POST("/inventories/restock", baseController.InventoryController.RestockInventory)
	r.DELETE("/inventories/:id", baseController.InventoryController.DeleteInventory)
	r.GET("/tenants", baseController.TenantController.GetTenants)
	r.GET("/tenants/:id", baseController.TenantController.GetTenantDetail)
	r.POST("/tenants", baseController.TenantController.CreateTenant)
	r.PUT("/tenants/:id", baseController.TenantController.UpdateTenant)
	r.DELETE("/tenants/:id", baseController.TenantController.DeleteTenant)
	r.GET("/orders", baseController.OrderController.GetOrders)
	r.GET("/orders/:id", baseController.OrderController.GetOrderDetail)
	r.GET("/users/:user_id/orders", baseController.OrderController.GetUserOrders)
	r.POST("/orders", baseController.OrderController.CreateOrder)
	r.POST("/orders/:id/cancel", baseController.OrderController.CancelOrder)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	return r
}
