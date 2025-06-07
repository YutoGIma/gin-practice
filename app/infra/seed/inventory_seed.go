package seed

import (
	"gorm.io/gorm"
	"myapp/app/model"
)

func SeedInventories(db *gorm.DB) error {
	inventories := []model.Inventory{
		{
			ProductId: 1,
			Quantity:  100,
		},
		{
			ProductId: 2,
			Quantity:  50,
		},
		{
			ProductId: 3,
			Quantity:  75,
		},
	}

	for _, inventory := range inventories {
		if err := db.Create(&inventory).Error; err != nil {
			return err
		}
	}

	return nil
}
