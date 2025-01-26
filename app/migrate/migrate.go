package main

import (
	"myapp/app/infra"
)

func main() {
	db := infra.SetupDB()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			panic("Failed to get database instance")
		}
		sqlDB.Close()
	}()

	infra.DBMigration(db)
}
