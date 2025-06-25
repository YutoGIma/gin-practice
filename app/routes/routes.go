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
	
	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	// User endpoints
	userGroup := r.Group("/users")
	{
		userGroup.GET("", baseController.UserController.GetUsers)
		userGroup.GET("/:id", baseController.UserController.GetUserDetail)
		userGroup.POST("", baseController.UserController.CreateUser)
		userGroup.PUT("/:id", baseController.UserController.UpdateUser)
		userGroup.DELETE("/:id", baseController.UserController.DeleteUser)
		userGroup.GET("/:user_id/orders", baseController.OrderController.GetUserOrders)
	}
	
	// Product endpoints
	productGroup := r.Group("/products")
	{
		productGroup.GET("", baseController.ProductController.GetProducts)
		productGroup.POST("", baseController.ProductController.CreateProduct)
		productGroup.PUT("/:id", baseController.ProductController.UpdateProduct)
		productGroup.DELETE("/:id", baseController.ProductController.DeleteProduct)
	}
	
	// Inventory endpoints
	inventoryGroup := r.Group("/inventories")
	{
		inventoryGroup.GET("", baseController.InventoryController.GetInventories)
		inventoryGroup.POST("", baseController.InventoryController.CreateInventory)
		inventoryGroup.PUT("/:id", baseController.InventoryController.UpdateInventory)
		inventoryGroup.DELETE("/:id", baseController.InventoryController.DeleteInventory)
		inventoryGroup.POST("/restock", baseController.InventoryController.RestockInventory)
		
		// Price setting endpoints
		inventoryGroup.POST("/:id/prices", baseController.PriceSettingController.CreatePriceSetting)
		inventoryGroup.GET("/:id/prices", baseController.PriceSettingController.GetPriceSettingsByInventoryID)
		inventoryGroup.GET("/:id/prices/current", baseController.PriceSettingController.GetCurrentPriceSetting)
		inventoryGroup.PUT("/:id/prices/:price_id", baseController.PriceSettingController.UpdatePriceSetting)
		inventoryGroup.DELETE("/:id/prices/:price_id", baseController.PriceSettingController.DeletePriceSetting)
	}
	
	// Tenant endpoints
	tenantGroup := r.Group("/tenants")
	{
		tenantGroup.GET("", baseController.TenantController.GetTenants)
		tenantGroup.GET("/:id", baseController.TenantController.GetTenantDetail)
		tenantGroup.POST("", baseController.TenantController.CreateTenant)
		tenantGroup.PUT("/:id", baseController.TenantController.UpdateTenant)
		tenantGroup.DELETE("/:id", baseController.TenantController.DeleteTenant)
	}
	
	// Order endpoints
	orderGroup := r.Group("/orders")
	{
		orderGroup.GET("", baseController.OrderController.GetOrders)
		orderGroup.GET("/:id", baseController.OrderController.GetOrderDetail)
		orderGroup.POST("", baseController.OrderController.CreateOrder)
		orderGroup.POST("/:id/cancel", baseController.OrderController.CancelOrder)
	}
	
	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	return r
}
