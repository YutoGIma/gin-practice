package usecase

import (
	"myapp/app/errors"
	"myapp/app/model"
	"myapp/app/service"
)

type TenantUseCase struct {
	tenantService service.TenantService
}

func NewTenantUseCase(tenantService service.TenantService) TenantUseCase {
	return TenantUseCase{
		tenantService: tenantService,
	}
}

func (uc *TenantUseCase) Create(input model.Tenant) (*model.Tenant, error) {
	tenant, err := uc.tenantService.CreateTenant(input)
	if err != nil {
		return nil, errors.NewInternalError("Failed to create tenant", err)
	}
	return tenant, nil
}

func (uc *TenantUseCase) Update(id uint, input model.Tenant) (*model.Tenant, error) {
	_, err := uc.tenantService.GetTenantDetail(id)
	if err != nil {
		return nil, errors.NewNotFoundError("Tenant not found")
	}

	updatedTenant, err := uc.tenantService.UpdateTenant(id, input)
	if err != nil {
		return nil, errors.NewInternalError("Failed to update tenant", err)
	}

	return updatedTenant, nil
}

func (uc *TenantUseCase) Delete(id uint) error {
	if err := uc.tenantService.DeleteTenant(id); err != nil {
		return errors.NewInternalError("Failed to delete tenant", err)
	}
	return nil
}

func (uc *TenantUseCase) List() ([]model.Tenant, error) {
	tenants, err := uc.tenantService.GetTenants()
	if err != nil {
		return nil, errors.NewInternalError("Failed to list tenants", err)
	}
	return tenants, nil
}
