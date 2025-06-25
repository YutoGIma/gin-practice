package infra

import (
	"myapp/app/infra/seed"
	"myapp/app/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	dsn := "host=db user=user password=password dbname=myapp port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	return db
}

func DBMigration(db *gorm.DB) {
	// 重複するバーコードを持つ商品を削除（ユニークインデックスを作成する前に）
	if err := cleanupDuplicateBarcodes(db); err != nil {
		panic("Failed to cleanup duplicate barcodes: " + err.Error())
	}

	if err := db.AutoMigrate(model.GetModels()...); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	// シードデータの追加
	if err := seed.SeedData(db); err != nil {
		panic("Failed to seed data: " + err.Error())
	}
}

// cleanupDuplicateBarcodes は重複するバーコードを持つ商品を削除します
func cleanupDuplicateBarcodes(db *gorm.DB) error {
	// 重複するバーコードを特定して古いレコードを削除
	query := `
		DELETE FROM products 
		WHERE id NOT IN (
			SELECT MIN(id) 
			FROM products 
			GROUP BY barcode
		)`
	
	return db.Exec(query).Error
}
