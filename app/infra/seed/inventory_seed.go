package seed

import (
	"myapp/app/model"

	"gorm.io/gorm"
)

func SeedInventories(db *gorm.DB) error {
	inventories := []model.Inventory{
		{
			ProductID: 1,
			TenantID:  1, // 本社
			Quantity:  100,
		},
		{
			ProductID: 2,
			TenantID:  1, // 本社
			Quantity:  50,
		},
		{
			ProductID: 3,
			TenantID:  2, // 大阪支店
			Quantity:  75,
		},
	}

	for _, inventory := range inventories {
		// Upsert: ProductIDとTenantIDの組み合わせをキーとして存在する場合は更新、存在しない場合は作成
		if err := db.Where("product_id = ? AND tenant_id = ?", inventory.ProductID, inventory.TenantID).
			Assign(inventory).
			FirstOrCreate(&inventory).Error; err != nil {
			return err
		}
	}

	return nil
}
