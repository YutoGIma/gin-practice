package main

import (
	"myapp/app/controller"
	"myapp/app/infra"
	"myapp/app/routes"
)

func main() {
	db := infra.SetupDB()
	db.AutoMigrate()

	baseController := controller.NewBaseController(db)

	// Use the db variable to avoid the "declared and not used" error
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get database instance")
	}
	defer sqlDB.Close()

	r := routes.SetupRouter(baseController)
	r.Run(":8080")
}
