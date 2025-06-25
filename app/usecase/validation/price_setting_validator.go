package validation

import (
	"myapp/app/errors"
	"myapp/app/usecase/request"
	"time"
)

// PriceSettingValidator handles price setting-related validations
type PriceSettingValidator struct {
	BaseValidator
}

// NewPriceSettingValidator creates a new price setting validator
func NewPriceSettingValidator() PriceSettingValidator {
	return PriceSettingValidator{
		BaseValidator: NewBaseValidator(),
	}
}

// ValidateCreatePriceSettingRequest validates price setting creation request
func (v PriceSettingValidator) ValidateCreatePriceSettingRequest(req request.CreatePriceSettingRequest) error {
	if err := v.ValidateID(req.InventoryID, "在庫ID"); err != nil {
		return err
	}
	
	if req.Price <= 0 {
		return errors.NewValidationError("価格は0より大きい値である必要があります")
	}
	
	if req.SalePrice != nil && *req.SalePrice <= 0 {
		return errors.NewValidationError("セール価格は0より大きい値である必要があります")
	}
	
	if req.SalePrice != nil && *req.SalePrice >= req.Price {
		return errors.NewValidationError("セール価格は通常価格より安く設定してください")
	}
	
	if err := v.validateDateRange(req.StartDate, req.EndDate); err != nil {
		return err
	}
	
	return nil
}

// ValidateUpdatePriceSettingRequest validates price setting update request
func (v PriceSettingValidator) ValidateUpdatePriceSettingRequest(req request.UpdatePriceSettingRequest) error {
	if req.Price != nil && *req.Price <= 0 {
		return errors.NewValidationError("価格は0より大きい値である必要があります")
	}
	
	if req.SalePrice != nil && *req.SalePrice != nil && **req.SalePrice <= 0 {
		return errors.NewValidationError("セール価格は0より大きい値である必要があります")
	}
	
	if req.Price != nil && req.SalePrice != nil && *req.SalePrice != nil {
		if **req.SalePrice >= *req.Price {
			return errors.NewValidationError("セール価格は通常価格より安く設定してください")
		}
	}
	
	if req.StartDate != nil && req.EndDate != nil {
		if err := v.validateDateRange(*req.StartDate, *req.EndDate); err != nil {
			return err
		}
	}
	
	return nil
}

// validateDateRange validates that start date is before end date
func (v PriceSettingValidator) validateDateRange(startDate time.Time, endDate *time.Time) error {
	now := time.Now()
	
	// Start date should not be too far in the past (more than 1 day)
	if startDate.Before(now.AddDate(0, 0, -1)) {
		return errors.NewValidationError("開始日は過去1日以内で設定してください")
	}
	
	if endDate != nil {
		if endDate.Before(startDate) {
			return errors.NewValidationError("終了日は開始日以降で設定してください")
		}
		
		if endDate.Before(now) {
			return errors.NewValidationError("終了日は現在日時以降で設定してください")
		}
		
		// End date should not be more than 1 year in the future
		if endDate.After(now.AddDate(1, 0, 0)) {
			return errors.NewValidationError("終了日は1年以内で設定してください")
		}
	}
	
	return nil
}

// ValidatePriceSettingPeriod validates that the period doesn't overlap with existing active price settings
func (v PriceSettingValidator) ValidatePriceSettingPeriod(hasOverlap bool) error {
	if hasOverlap {
		return errors.NewValidationError("指定した期間は既存の価格設定と重複しています")
	}
	return nil
}