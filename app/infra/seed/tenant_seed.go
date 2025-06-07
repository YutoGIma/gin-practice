package seed

import (
	"gorm.io/gorm"
	"myapp/app/model"
	"time"
)

func SeedTenants(db *gorm.DB) error {
	tenants := []model.Tenant{
		{
			Name:        "本社",
			Code:        "HQ001",
			Address:     "東京都千代田区丸の内1-1-1",
			PhoneNumber: "03-1234-5678",
			Email:       "hq@example.com",
			IsActive:    true,
			OpenedAt:    time.Now(),
		},
		{
			Name:        "大阪支店",
			Code:        "OSK001",
			Address:     "大阪府大阪市中央区本町1-1-1",
			PhoneNumber: "06-1234-5678",
			Email:       "osaka@example.com",
			IsActive:    true,
			OpenedAt:    time.Now(),
		},
		{
			Name:        "名古屋支店",
			Code:        "NGY001",
			Address:     "愛知県名古屋市中区栄1-1-1",
			PhoneNumber: "052-123-4567",
			Email:       "nagoya@example.com",
			IsActive:    true,
			OpenedAt:    time.Now(),
		},
	}

	for _, tenant := range tenants {
		if err := db.Create(&tenant).Error; err != nil {
			return err
		}
	}

	return nil
}
