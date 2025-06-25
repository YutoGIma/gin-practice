package seed

import (
	"myapp/app/model"

	"gorm.io/gorm"
)

func SeedInventories(db *gorm.DB) error {
	inventories := []model.Inventory{
		{
			ProductID: 1,
			Quantity:  100,
		},
		{
			ProductID: 2,
			Quantity:  50,
		},
		{
			ProductID: 3,
			Quantity:  75,
		},
	}

	for _, inventory := range inventories {
		// 既存の在庫をProductIDでチェック
		var existingInventory model.Inventory
		if err := db.Where("product_id = ?", inventory.ProductID).First(&existingInventory).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// 存在しない場合は作成
				if err := db.Create(&inventory).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
		// 既に存在する場合はスキップ
	}

	return nil
}
