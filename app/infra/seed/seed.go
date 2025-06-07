package seed

import (
	"gorm.io/gorm"
)

// SeedData は全てのシードデータを実行します
func SeedData(db *gorm.DB) error {
	// シードデータの実行順序を制御
	seeders := []func(*gorm.DB) error{
		SeedUsers,       // ユーザーを最初に作成
		SeedProducts,    // 次に商品を作成
		SeedInventories, // 最後に在庫を作成
	}

	// 各シーダーを順番に実行
	for _, seeder := range seeders {
		if err := seeder(db); err != nil {
			return err
		}
	}

	return nil
}
