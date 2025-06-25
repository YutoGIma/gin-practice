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
		// 既存の商品をバーコードでチェック
		var existingProduct model.Product
		if err := db.Where("barcode = ?", product.Barcode).First(&existingProduct).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// 存在しない場合は作成
				if err := db.Create(&product).Error; err != nil {
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
