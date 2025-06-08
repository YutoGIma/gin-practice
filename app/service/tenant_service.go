package service

import (
	"myapp/app/model"

	"gorm.io/gorm"
)

type TenantService struct {
	DB *gorm.DB
}

func NewTenantService(db *gorm.DB) TenantService {
	return TenantService{
		DB: db,
	}
}

func (s *TenantService) GetTenants() ([]model.Tenant, error) {
	var tenants []model.Tenant
	if err := s.DB.Find(&tenants).Error; err != nil {
		return nil, err
	}
	return tenants, nil
}

func (s *TenantService) GetTenantDetail(id uint) (*model.Tenant, error) {
	var tenant model.Tenant
	if err := s.DB.First(&tenant, id).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (s *TenantService) CreateTenant(tenant model.Tenant) (*model.Tenant, error) {
	if err := s.DB.Create(&tenant).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (s *TenantService) UpdateTenant(id uint, tenant model.Tenant) (*model.Tenant, error) {
	if err := s.DB.Model(&model.Tenant{}).Where("id = ?", id).Updates(tenant).Error; err != nil {
		return nil, err
	}
	return s.GetTenantDetail(id)
}

func (s *TenantService) DeleteTenant(id uint) error {
	return s.DB.Delete(&model.Tenant{}, id).Error
}
