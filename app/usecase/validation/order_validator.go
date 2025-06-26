package validation

import (
	"fmt"
	"myapp/app/errors"
	"myapp/app/model"
	"myapp/app/usecase/request"
)

// OrderValidator handles order-related validations
type OrderValidator struct {
	BaseValidator
}

// NewOrderValidator creates a new order validator
func NewOrderValidator() OrderValidator {
	return OrderValidator{
		BaseValidator: NewBaseValidator(),
	}
}

// ValidateCreateOrderRequest validates order creation request
func (v OrderValidator) ValidateCreateOrderRequest(req request.CreateOrderRequest) error {
	if err := v.ValidateID(req.UserID, "ユーザーID"); err != nil {
		return err
	}

	if err := v.ValidateID(req.TenantID, "テナントID"); err != nil {
		return err
	}

	if len(req.OrderItems) == 0 {
		return errors.NewValidationError("注文明細は1件以上必要です")
	}

	// Validate each order item
	for i, item := range req.OrderItems {
		if err := v.validateOrderItem(item, i+1); err != nil {
			return err
		}
	}

	return nil
}

// validateOrderItem validates a single order item
func (v OrderValidator) validateOrderItem(item request.OrderItemRequest, index int) error {
	if err := v.ValidateID(item.ProductID, fmt.Sprintf("注文明細%dの商品ID", index)); err != nil {
		return err
	}

	if err := v.ValidatePositiveNumber(item.Quantity, fmt.Sprintf("注文明細%dの数量", index)); err != nil {
		return err
	}

	return nil
}

// ValidateCancelOrder validates order cancellation
func (v OrderValidator) ValidateCancelOrder(order *model.Order) error {
	if order == nil {
		return errors.NewNotFoundError("注文が見つかりません")
	}

	if order.Status != model.OrderStatusPending {
		return errors.NewValidationError("この注文はキャンセルできません。ステータス: " + string(order.Status))
	}

	return nil
}

// ValidateSufficientInventory validates inventory for order item
func (v OrderValidator) ValidateSufficientInventory(availableQuantity, requestedQuantity int, productName string) error {
	if availableQuantity < requestedQuantity {
		return errors.NewValidationError(fmt.Sprintf("商品「%s」の在庫が不足しています。現在の在庫: %d", productName, availableQuantity))
	}
	return nil
}
