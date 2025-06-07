package usecase

import (
	"gorm.io/gorm"
	"myapp/app/errors"
	"myapp/app/model"
)

type TenantUseCase struct {
	db *gorm.DB
}

func NewTenantUseCase(db *gorm.DB) *TenantUseCase {
	return &TenantUseCase{
		db: db,
	}
}

func (uc *TenantUseCase) Create(input model.Tenant) (*model.Tenant, error) {
	if err := uc.db.Create(&input).Error; err != nil {
		return nil, errors.NewInternalError("Failed to create tenant", err)
	}
	return &input, nil
}

func (uc *TenantUseCase) Update(id uint, input model.Tenant) (*model.Tenant, error) {
	var tenant model.Tenant
	if err := uc.db.First(&tenant, id).Error; err != nil {
		return nil, errors.NewNotFoundError("Tenant not found")
	}

	if err := uc.db.Model(&tenant).Updates(input).Error; err != nil {
		return nil, errors.NewInternalError("Failed to update tenant", err)
	}

	return &tenant, nil
}

func (uc *TenantUseCase) Delete(id uint) error {
	if err := uc.db.Delete(&model.Tenant{}, id).Error; err != nil {
		return errors.NewInternalError("Failed to delete tenant", err)
	}
	return nil
}

func (uc *TenantUseCase) GetByID(id uint) (*model.Tenant, error) {
	var tenant model.Tenant
	if err := uc.db.First(&tenant, id).Error; err != nil {
		return nil, errors.NewNotFoundError("Tenant not found")
	}
	return &tenant, nil
}

func (uc *TenantUseCase) List() ([]model.Tenant, error) {
	var tenants []model.Tenant
	if err := uc.db.Find(&tenants).Error; err != nil {
		return nil, errors.NewInternalError("Failed to list tenants", err)
	}
	return tenants, nil
}
