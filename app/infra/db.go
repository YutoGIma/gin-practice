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
	if err := db.AutoMigrate(model.GetModels()...); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	// シードデータの追加
	if err := seed.SeedData(db); err != nil {
		panic("Failed to seed data: " + err.Error())
	}
}
