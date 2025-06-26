package usecase

import (
	"fmt"
	"myapp/app/errors"
	"myapp/app/model"
	"myapp/app/service"
	"myapp/app/usecase/request"
	"myapp/app/usecase/validation"
	"time"

	"gorm.io/gorm"
)

type OrderUseCase struct {
	orderService     service.OrderService
	inventoryService service.InventoryService
	productService   service.ProductService
	validator        validation.OrderValidator
}

func NewOrderUseCase(orderService service.OrderService, inventoryService service.InventoryService, productService service.ProductService) OrderUseCase {
	return OrderUseCase{
		orderService:     orderService,
		inventoryService: inventoryService,
		productService:   productService,
		validator:        validation.NewOrderValidator(),
	}
}

func (uc *OrderUseCase) GetOrders() ([]model.Order, error) {
	return uc.orderService.GetOrders()
}

func (uc *OrderUseCase) GetOrderByID(id uint) (*model.Order, error) {
	order, err := uc.orderService.GetOrderByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("注文が見つかりません")
		}
		return nil, errors.NewInternalError("注文の取得に失敗しました", err)
	}
	return order, nil
}

func (uc *OrderUseCase) GetOrdersByUserID(userID uint) ([]model.Order, error) {
	return uc.orderService.GetOrdersByUserID(userID)
}

func (uc *OrderUseCase) CreateOrder(req request.CreateOrderRequest) (*model.Order, error) {
	// リクエストのバリデーション
	if err := uc.validator.ValidateCreateOrderRequest(req); err != nil {
		return nil, err
	}

	// トランザクション開始
	tx := uc.orderService.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 注文の初期化
	order := &model.Order{
		UserID:     req.UserID,
		TenantID:   req.TenantID,
		Status:     model.OrderStatusPending,
		OrderedAt:  time.Now(),
		OrderItems: []model.OrderItem{},
	}

	var totalAmount float64

	// 各商品の在庫確認と注文明細の作成
	for _, item := range req.OrderItems {
		// 商品情報の取得
		product, err := uc.productService.GetProductByID(item.ProductID)
		if err != nil {
			tx.Rollback()
			if err == gorm.ErrRecordNotFound {
				return nil, errors.NewNotFoundError(fmt.Sprintf("商品ID %d が見つかりません", item.ProductID))
			}
			return nil, errors.NewInternalError("商品情報の取得に失敗しました", err)
		}

		// 在庫の確認
		inventory, err := uc.inventoryService.GetInventoryByProductAndTenant(item.ProductID, req.TenantID)
		if err != nil {
			tx.Rollback()
			if err == gorm.ErrRecordNotFound {
				return nil, errors.NewNotFoundError(fmt.Sprintf("商品ID %d の在庫情報が見つかりません", item.ProductID))
			}
			return nil, errors.NewInternalError("在庫情報の取得に失敗しました", err)
		}

		// 在庫数の確認
		if err := uc.validator.ValidateSufficientInventory(inventory.Quantity, item.Quantity, product.Name); err != nil {
			tx.Rollback()
			return nil, err
		}

		// 在庫の更新
		inventory.Quantity -= item.Quantity
		if err := uc.inventoryService.UpdateInventoryWithTx(tx, inventory); err != nil {
			tx.Rollback()
			return nil, errors.NewInternalError("在庫の更新に失敗しました", err)
		}

		// 注文明細の作成
		subtotal := product.PurchasePrice * float64(item.Quantity)
		orderItem := model.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.PurchasePrice,
			Subtotal:  subtotal,
		}
		order.OrderItems = append(order.OrderItems, orderItem)
		totalAmount += subtotal
	}

	order.TotalAmount = totalAmount

	// 注文の作成
	if err := uc.orderService.CreateOrderWithTx(tx, order); err != nil {
		tx.Rollback()
		return nil, errors.NewInternalError("注文の作成に失敗しました", err)
	}

	// 注文ステータスを完了に更新
	order.Status = model.OrderStatusCompleted
	if err := tx.Model(order).Update("status", model.OrderStatusCompleted).Error; err != nil {
		tx.Rollback()
		return nil, errors.NewInternalError("注文ステータスの更新に失敗しました", err)
	}

	// トランザクションのコミット
	if err := tx.Commit().Error; err != nil {
		return nil, errors.NewInternalError("トランザクションのコミットに失敗しました", err)
	}

	// 作成した注文を再取得（関連データを含む）
	createdOrder, err := uc.orderService.GetOrderByID(order.ID)
	if err != nil {
		return nil, errors.NewInternalError("作成した注文の取得に失敗しました", err)
	}

	return createdOrder, nil
}

func (uc *OrderUseCase) CancelOrder(id uint) error {
	order, err := uc.orderService.GetOrderByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("注文が見つかりません")
		}
		return errors.NewInternalError("注文の取得に失敗しました", err)
	}

	// 注文キャンセルのバリデーション
	if err := uc.validator.ValidateCancelOrder(order); err != nil {
		return err
	}

	// TODO: 在庫の復元処理を実装

	return uc.orderService.UpdateOrderStatus(id, model.OrderStatusCancelled)
}
