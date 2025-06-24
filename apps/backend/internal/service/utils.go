package service

import (
	"encoding/json"
	"fmt"

	"github.com/WBTechSchool/L0/internal/gen"
	"github.com/WBTechSchool/L0/internal/models"
)

func convertOrderToOrderResponse(order interface{}) (*models.OrderResponse, error) {
	if order, ok := order.(gen.GetOrderByIDRow); ok {
		var items []models.OrderItem
		if order.Items != nil {
			if err := json.Unmarshal(order.Items, &items); err != nil {
				return nil, fmt.Errorf("failed to unmarshal items: %w", err)
			}
		}

		response := &models.OrderResponse{
			OrderID:           order.OrderID,
			CustomerID:        order.CustomerID,
			TrackNumber:       order.TrackNumber,
			Locale:            order.Locale,
			InternalSignature: order.InternalSignature,
			DeliveryService:   order.DeliveryService,
			Entry:             order.Entry,
			ShardKey:          order.ShardKey,
			OffShard:          order.OffShard,
			SmID:              order.SmID,

			Delivery: models.Delivery{
				Name:    *order.Name,
				Phone:   *order.Phone,
				Zip:     *order.Zip,
				City:    *order.City,
				Address: *order.Address,
				Region:  *order.Region,
				Email:   *order.Email,
			},

			Payments: models.Payment{
				Transaction:  *order.Transaction,
				RequestID:    *order.RequestID,
				Currency:     *order.Currency,
				Provider:     *order.Provider,
				Amount:       *order.Amount,
				PaymentDt:    *order.PaymentDt,
				Bank:         *order.Bank,
				DeliveryCost: *order.DeliveryCost,
				GoodsTotal:   *order.GoodsTotal,
				CustomFee:    *order.CustomFee,
			},

			Items: items,
		}

		return response, nil
	}

	return nil, fmt.Errorf("type mismatch")
}
