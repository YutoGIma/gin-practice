package usecase

import (
	"myapp/app/errors"
	"myapp/app/model"
	"myapp/app/service"
)

type ProductUseCase struct {
	productService service.ProductService
}

func NewProductUseCase(productService service.ProductService) ProductUseCase {
	return ProductUseCase{
		productService: productService,
	}
}

func (uc *ProductUseCase) Create(input model.Product) (*model.Product, error) {
	if err := model.ValidateCreateProduct(input); err != nil {
		return nil, errors.NewValidationError(err.Error())
	}

	if err := uc.productService.CreateProduct(input); err != nil {
		return nil, errors.NewInternalError("Failed to create product", err)
	}

	return &input, nil
}

func (uc *ProductUseCase) Update(id uint, input model.Product) (*model.Product, error) {
	product, err := uc.productService.GetProductDetail(id)
	if err != nil {
		return nil, errors.NewNotFoundError("Product not found")
	}

	if err := model.ValidateCreateProduct(input); err != nil {
		return nil, errors.NewValidationError(err.Error())
	}

	if err := uc.productService.UpdateProduct(input); err != nil {
		return nil, errors.NewInternalError("Failed to update product", err)
	}

	return &product, nil
}

func (uc *ProductUseCase) Delete(id uint) error {
	if err := uc.productService.DeleteProduct(id); err != nil {
		return errors.NewInternalError("Failed to delete product", err)
	}
	return nil
}

func (uc *ProductUseCase) GetByID(id uint) (*model.Product, error) {
	product, err := uc.productService.GetProductDetail(id)
	if err != nil {
		return nil, errors.NewNotFoundError("Product not found")
	}
	return &product, nil
}

func (uc *ProductUseCase) List() ([]model.Product, error) {
	products, err := uc.productService.GetProducts()
	if err != nil {
		return nil, errors.NewInternalError("Failed to list products", err)
	}
	return products, nil
}
