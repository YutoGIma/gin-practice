package usecase

import (
	"myapp/app/errors"
	"myapp/app/model"
	"myapp/app/service"
	"myapp/app/usecase/request"
	"myapp/app/usecase/validation"

	"gorm.io/gorm"
)

type PriceSettingUseCase struct {
	priceSettingService service.PriceSettingService
	inventoryService    service.InventoryService
	validator           validation.PriceSettingValidator
}

func NewPriceSettingUseCase(priceSettingService service.PriceSettingService, inventoryService service.InventoryService) PriceSettingUseCase {
	return PriceSettingUseCase{
		priceSettingService: priceSettingService,
		inventoryService:    inventoryService,
		validator:           validation.NewPriceSettingValidator(),
	}
}

func (uc *PriceSettingUseCase) GetPriceSettingsByInventoryID(inventoryID uint) ([]model.PriceSetting, error) {
	// 在庫の存在確認
	_, err := uc.inventoryService.GetInventoryDetail(inventoryID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("在庫が見つかりません")
		}
		return nil, errors.NewInternalError("在庫の取得に失敗しました", err)
	}

	priceSettings, err := uc.priceSettingService.GetPriceSettingsByInventoryID(inventoryID)
	if err != nil {
		return nil, errors.NewInternalError("価格設定の取得に失敗しました", err)
	}

	return priceSettings, nil
}

func (uc *PriceSettingUseCase) GetCurrentPriceSetting(inventoryID uint) (*model.PriceSetting, error) {
	// 在庫の存在確認
	_, err := uc.inventoryService.GetInventoryDetail(inventoryID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("在庫が見つかりません")
		}
		return nil, errors.NewInternalError("在庫の取得に失敗しました", err)
	}

	priceSetting, err := uc.priceSettingService.GetCurrentPriceSetting(inventoryID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("有効な価格設定が見つかりません")
		}
		return nil, errors.NewInternalError("現在の価格設定の取得に失敗しました", err)
	}

	return priceSetting, nil
}

func (uc *PriceSettingUseCase) GetPriceSettingByID(id uint) (*model.PriceSetting, error) {
	priceSetting, err := uc.priceSettingService.GetPriceSettingByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("価格設定が見つかりません")
		}
		return nil, errors.NewInternalError("価格設定の取得に失敗しました", err)
	}

	return priceSetting, nil
}

func (uc *PriceSettingUseCase) CreatePriceSetting(req request.CreatePriceSettingRequest) (*model.PriceSetting, error) {
	// リクエストのバリデーション
	if err := uc.validator.ValidateCreatePriceSettingRequest(req); err != nil {
		return nil, err
	}

	// 在庫の存在確認
	_, err := uc.inventoryService.GetInventoryDetail(req.InventoryID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("在庫が見つかりません")
		}
		return nil, errors.NewInternalError("在庫の取得に失敗しました", err)
	}

	// 期間の重複チェック
	hasOverlap, err := uc.priceSettingService.CheckOverlappingPeriods(req.InventoryID, req.StartDate, req.EndDate, nil)
	if err != nil {
		return nil, errors.NewInternalError("期間の重複チェックに失敗しました", err)
	}

	if err := uc.validator.ValidatePriceSettingPeriod(hasOverlap); err != nil {
		return nil, err
	}

	// トランザクション開始
	tx := uc.priceSettingService.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 価格設定を作成
	priceSetting := &model.PriceSetting{
		InventoryID: req.InventoryID,
		Price:       req.Price,
		SalePrice:   req.SalePrice,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		IsActive:    true,
		Note:        req.Note,
	}

	if err := uc.priceSettingService.CreatePriceSettingWithTx(tx, priceSetting); err != nil {
		tx.Rollback()
		return nil, errors.NewInternalError("価格設定の作成に失敗しました", err)
	}

	// 他の有効な価格設定を無効化
	if err := uc.priceSettingService.DeactivateOldPriceSettings(req.InventoryID, priceSetting.ID); err != nil {
		tx.Rollback()
		return nil, errors.NewInternalError("既存価格設定の無効化に失敗しました", err)
	}

	// トランザクションのコミット
	if err := tx.Commit().Error; err != nil {
		return nil, errors.NewInternalError("トランザクションのコミットに失敗しました", err)
	}

	// 作成した価格設定を再取得
	createdPriceSetting, err := uc.priceSettingService.GetPriceSettingByID(priceSetting.ID)
	if err != nil {
		return nil, errors.NewInternalError("作成した価格設定の取得に失敗しました", err)
	}

	return createdPriceSetting, nil
}

func (uc *PriceSettingUseCase) UpdatePriceSetting(id uint, req request.UpdatePriceSettingRequest) (*model.PriceSetting, error) {
	// リクエストのバリデーション
	if err := uc.validator.ValidateUpdatePriceSettingRequest(req); err != nil {
		return nil, err
	}

	// 既存の価格設定を取得
	existingPriceSetting, err := uc.priceSettingService.GetPriceSettingByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("価格設定が見つかりません")
		}
		return nil, errors.NewInternalError("価格設定の取得に失敗しました", err)
	}

	// 更新内容を適用
	if req.Price != nil {
		existingPriceSetting.Price = *req.Price
	}
	if req.SalePrice != nil {
		existingPriceSetting.SalePrice = *req.SalePrice
	}
	if req.StartDate != nil {
		existingPriceSetting.StartDate = *req.StartDate
	}
	if req.EndDate != nil {
		existingPriceSetting.EndDate = *req.EndDate
	}
	if req.IsActive != nil {
		existingPriceSetting.IsActive = *req.IsActive
	}
	if req.Note != nil {
		existingPriceSetting.Note = *req.Note
	}

	// 更新
	if err := uc.priceSettingService.UpdatePriceSetting(existingPriceSetting); err != nil {
		return nil, errors.NewInternalError("価格設定の更新に失敗しました", err)
	}

	// 更新した価格設定を再取得
	updatedPriceSetting, err := uc.priceSettingService.GetPriceSettingByID(id)
	if err != nil {
		return nil, errors.NewInternalError("更新した価格設定の取得に失敗しました", err)
	}

	return updatedPriceSetting, nil
}

func (uc *PriceSettingUseCase) DeletePriceSetting(id uint) error {
	// 価格設定の存在確認
	_, err := uc.priceSettingService.GetPriceSettingByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("価格設定が見つかりません")
		}
		return errors.NewInternalError("価格設定の取得に失敗しました", err)
	}

	// 削除
	if err := uc.priceSettingService.DeletePriceSetting(id); err != nil {
		return errors.NewInternalError("価格設定の削除に失敗しました", err)
	}

	return nil
}