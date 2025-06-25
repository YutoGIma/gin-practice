package seed

import (
	"gorm.io/gorm"
)

// SeedData は全てのシードデータを実行します
func SeedData(db *gorm.DB) error {
	// シードデータの実行順序を制御
	seeders := []func(*gorm.DB) error{
		SeedTenants,
		SeedUsers,
		SeedProducts,
		SeedInventories,
		SeedPriceSettings, // 在庫作成後に価格設定を作成
	}

	// 各シーダーを順番に実行
	for _, seeder := range seeders {
		if err := seeder(db); err != nil {
			return err
		}
	}

	return nil
}
