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
		if err := db.Create(&inventory).Error; err != nil {
			return err
		}
	}

	return nil
}
