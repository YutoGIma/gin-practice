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
		// Upsert: メールアドレスをキーとして存在する場合は更新、存在しない場合は作成
		if err := db.Where("email = ?", user.Email).
			Assign(user).
			FirstOrCreate(&user).Error; err != nil {
			return err
		}
	}

	return nil
}
