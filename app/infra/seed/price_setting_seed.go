package seed

import (
	"myapp/app/model"
	"time"

	"gorm.io/gorm"
)

func SeedPriceSettings(db *gorm.DB) error {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	nextWeek := now.AddDate(0, 0, 7)
	nextMonth := now.AddDate(0, 1, 0)
	
	// セール価格用のポインタ作成
	salePrice1 := 800.0
	salePrice2 := 1200.0
	
	// 終了日用のポインタ作成
	endDate1 := nextWeek
	endDate2 := nextMonth

	priceSettings := []model.PriceSetting{
		{
			InventoryID: 1, // 商品1の在庫
			Price:       1000.0,
			SalePrice:   &salePrice1, // セール価格: 800円
			StartDate:   now,
			EndDate:     &endDate1, // 1週間後まで
			IsActive:    true,
			Note:        "新商品キャンペーン価格",
		},
		{
			InventoryID: 2, // 商品2の在庫
			Price:       1500.0,
			SalePrice:   &salePrice2, // セール価格: 1200円
			StartDate:   tomorrow,
			EndDate:     &endDate2, // 1ヶ月後まで
			IsActive:    true,
			Note:        "春の特別セール",
		},
		{
			InventoryID: 3, // 商品3の在庫
			Price:       2000.0,
			SalePrice:   nil, // セール価格なし
			StartDate:   now,
			EndDate:     nil, // 終了日なし（無期限）
			IsActive:    true,
			Note:        "通常価格",
		},
	}

	for _, priceSetting := range priceSettings {
		// Upsert: InventoryIDをキーとして存在する場合は更新、存在しない場合は作成
		// 既存の価格設定を非アクティブにする
		db.Model(&model.PriceSetting{}).
			Where("inventory_id = ?", priceSetting.InventoryID).
			Update("is_active", false)
		
		// 新しい価格設定を作成
		if err := db.Create(&priceSetting).Error; err != nil {
			return err
		}
	}

	return nil
}