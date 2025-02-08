package controller

import (
	"myapp/app/service"

	"gorm.io/gorm"
)

type BaseController struct {
	UserController      UserController
	InventoryController InventoryController
	ProductController   ProductController
	// 他のコントローラーを追加する場合は、ここにフィールドを追加します
	// Example:
	// ProductController *ProductController
}

func NewBaseController(db *gorm.DB) *BaseController {
	baseService := service.NewBaseService(db)
	return &BaseController{
		UserController:      UserController{UserService: baseService.UserService},
		InventoryController: InventoryController{InventoryService: baseService.InventoryService},
		ProductController:   ProductController{ProductService: baseService.ProductService},
		// 他のコントローラーの初期化を追加します
		// Example:
		// ProductController: &ProductController{ProductService: productService},
	}
}
