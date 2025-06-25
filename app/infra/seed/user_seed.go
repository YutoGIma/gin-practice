package seed

import (
	"gorm.io/gorm"
	"myapp/app/model"
)

func SeedUsers(db *gorm.DB) error {
	users := []model.User{
		{
			Name:     "管理者",
			Email:    "admin@example.com",
			Password: "admin123", // 実際の運用ではハッシュ化が必要
		},
		{
			Name:     "一般ユーザー",
			Email:    "user@example.com",
			Password: "user123", // 実際の運用ではハッシュ化が必要
		},
	}

	for _, user := range users {
		// 既存のユーザーをメールアドレスでチェック
		var existingUser model.User
		if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// 存在しない場合は作成
				if err := db.Create(&user).Error; err != nil {
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
