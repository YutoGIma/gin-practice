package service

import (
	"myapp/app/model"
	"time"

	"gorm.io/gorm"
)

type PriceSettingService struct {
	DB *gorm.DB
}

func NewPriceSettingService(db *gorm.DB) PriceSettingService {
	return PriceSettingService{
		DB: db,
	}
}

func (s *PriceSettingService) GetPriceSettingsByInventoryID(inventoryID uint) ([]model.PriceSetting, error) {
	var priceSettings []model.PriceSetting
	err := s.DB.Where("inventory_id = ?", inventoryID).
		Order("start_date DESC").
		Find(&priceSettings).Error
	return priceSettings, err
}

func (s *PriceSettingService) GetPriceSettingByID(id uint) (*model.PriceSetting, error) {
	var priceSetting model.PriceSetting
	err := s.DB.Preload("Inventory.Product").Preload("Inventory.Tenant").
		First(&priceSetting, id).Error
	if err != nil {
		return nil, err
	}
	return &priceSetting, nil
}

func (s *PriceSettingService) GetCurrentPriceSetting(inventoryID uint) (*model.PriceSetting, error) {
	var priceSetting model.PriceSetting
	now := time.Now()
	
	err := s.DB.Where("inventory_id = ? AND is_active = ? AND start_date <= ? AND (end_date IS NULL OR end_date >= ?)", 
		inventoryID, true, now, now).
		Order("start_date DESC").
		First(&priceSetting).Error
	
	if err != nil {
		return nil, err
	}
	return &priceSetting, nil
}

func (s *PriceSettingService) CreatePriceSetting(priceSetting *model.PriceSetting) error {
	return s.DB.Create(priceSetting).Error
}

func (s *PriceSettingService) UpdatePriceSetting(priceSetting *model.PriceSetting) error {
	return s.DB.Save(priceSetting).Error
}

func (s *PriceSettingService) DeletePriceSetting(id uint) error {
	return s.DB.Delete(&model.PriceSetting{}, id).Error
}

func (s *PriceSettingService) CheckOverlappingPeriods(inventoryID uint, startDate time.Time, endDate *time.Time, excludeID *uint) (bool, error) {
	query := s.DB.Where("inventory_id = ? AND is_active = ?", inventoryID, true)
	
	if excludeID != nil {
		query = query.Where("id != ?", *excludeID)
	}
	
	// Check for overlapping periods
	if endDate != nil {
		// Both start and end dates are specified
		query = query.Where("(start_date <= ? AND (end_date IS NULL OR end_date >= ?))", *endDate, startDate)
	} else {
		// Only start date is specified (infinite end)
		query = query.Where("(end_date IS NULL OR end_date >= ?)", startDate)
	}
	
	var count int64
	err := query.Model(&model.PriceSetting{}).Count(&count).Error
	return count > 0, err
}

func (s *PriceSettingService) DeactivateOldPriceSettings(inventoryID uint, excludeID uint) error {
	return s.DB.Model(&model.PriceSetting{}).
		Where("inventory_id = ? AND id != ? AND is_active = ?", inventoryID, excludeID, true).
		Update("is_active", false).Error
}

func (s *PriceSettingService) BeginTransaction() *gorm.DB {
	return s.DB.Begin()
}

func (s *PriceSettingService) CreatePriceSettingWithTx(tx *gorm.DB, priceSetting *model.PriceSetting) error {
	return tx.Create(priceSetting).Error
}