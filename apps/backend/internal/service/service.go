package service

import "github.com/WBTechSchool/L0/internal/db"

type Service struct {
	OrderService OrderService
}

func NewService(db *db.Storage) *Service {
	orderService := NewOrderService(db)

	return &Service{
		OrderService: orderService,
	}
}
