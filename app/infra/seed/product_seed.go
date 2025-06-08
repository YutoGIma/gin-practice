package seed

import (
	"myapp/app/model"

	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) error {
	products := []model.Product{
		{
			Name:          "サンプル商品1",
			Barcode:       "4901234567890",
			CategoryId:    1,
			PurchasePrice: 1000.00,
		},
		{
			Name:          "サンプル商品2",
			Barcode:       "4902345678901",
			CategoryId:    2,
			PurchasePrice: 2000.00,
		},
		{
			Name:          "サンプル商品3",
			Barcode:       "4903456789012",
			CategoryId:    1,
			PurchasePrice: 1500.00,
		},
	}

	for _, product := range products {
		if err := db.Create(&product).Error; err != nil {
			return err
		}
	}

	return nil
}
