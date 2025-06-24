package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/WBTechSchool/L0/internal/db"
	"github.com/WBTechSchool/L0/internal/gen"
	"github.com/WBTechSchool/L0/internal/models"
)

type OrderService interface {
	GetOrderByID(ctx context.Context, id int) (*models.OrderResponse, error)
	GetOrders(ctx context.Context)
	CreateOrder(ctx context.Context, params models.CreateOrderRequest) (int64, error)
}

type orderService struct {
	db          *db.Storage
	cacheOrders map[int64]models.OrderResponse
}

func NewOrderService(db *db.Storage) OrderService {
	return &orderService{
		db:          db,
		cacheOrders: make(map[int64]models.OrderResponse, 0),
	}
}

func (o *orderService) GetOrderByID(ctx context.Context, id int) (*models.OrderResponse, error) {
	if response, ok := o.cacheOrders[int64(id)]; ok {
		slog.Info("Return cached order", slog.Int64("id", int64(id)))
		return &response, nil
	}

	order, err := o.db.Queries.GetOrderByID(ctx, int64(id))
	if err != nil {
		return nil, fmt.Errorf("Failed to get order: %w", err)
	}

	response, err := convertOrderToOrderResponse(order)
	if err != nil {
		return nil, fmt.Errorf("Failed to convert order: %w", err)
	}

	o.cacheOrders[int64(id)] = *response

	return response, nil
}

func (o *orderService) GetOrders(ctx context.Context) {
	rows, err := o.db.Queries.GetOrders(ctx)
	if err != nil {
		return
	}

	for _, row := range rows {
		response, err := convertOrderToOrderResponse(row)
		if err != nil {
			continue
		}

		o.cacheOrders[response.OrderID] = *response
	}
}

func (o *orderService) CreateOrder(ctx context.Context, params models.CreateOrderRequest) (int64, error) {
	tx, err := o.db.Pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("Failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	q := o.db.Queries.WithTx(tx)

	orderID, err := q.CreateOrder(ctx, gen.CreateOrderParams(params.Order))
	if err != nil {
		return 0, fmt.Errorf("Failed to create order: %w", err)
	}

	paymentParams := gen.CreatePaymentParams{
		OrderID:      orderID,
		Transaction:  params.Payment.Transaction,
		RequestID:    params.Payment.RequestID,
		Currency:     params.Payment.Currency,
		Provider:     params.Payment.Provider,
		Amount:       params.Payment.Amount,
		PaymentDt:    params.Payment.PaymentDt,
		Bank:         params.Payment.Bank,
		DeliveryCost: params.Payment.DeliveryCost,
		GoodsTotal:   params.Payment.GoodsTotal,
		CustomFee:    params.Payment.CustomFee,
	}
	err = q.CreatePayment(ctx, paymentParams)
	if err != nil {
		return 0, fmt.Errorf("Failed to create payment: %w", err)
	}

	deliveryParams := gen.CreateDeliveryParams{
		OrderID: orderID,
		Name:    params.Delivery.Name,
		Phone:   params.Delivery.Phone,
		Zip:     params.Delivery.Zip,
		City:    params.Delivery.City,
		Address: params.Delivery.Address,
		Region:  params.Delivery.Region,
		Email:   params.Delivery.Email,
	}
	err = q.CreateDelivery(ctx, deliveryParams)
	if err != nil {
		return 0, fmt.Errorf("Failed to create delivery: %w", err)
	}

	for _, item := range params.OrderItems {
		err = q.CreateOrderItem(ctx, gen.CreateOrderItemParams{
			OrderID:     orderID,
			ChrtID:      item.ChrtID,
			TrackNumber: item.TrackNumber,
			Price:       item.Price,
			Rid:         item.Rid,
			Name:        item.Name,
			Sale:        item.Sale,
			Size:        item.Size,
			TotalPrice:  item.TotalPrice,
			NmID:        item.NmID,
			Brand:       item.Brand,
			Status:      item.Status,
		})
		if err != nil {
			return 0, fmt.Errorf("failed to create order item: %w", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return orderID, nil
}
