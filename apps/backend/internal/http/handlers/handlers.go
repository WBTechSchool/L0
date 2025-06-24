package handlers

import "github.com/WBTechSchool/L0/internal/service"

type Handlers struct {
	OrderHandlers OrderHandlers
}

func NewHandlers(s *service.Service) *Handlers {
	orderHandlers := NewOrderHandlers(s)

	return &Handlers{
		OrderHandlers: orderHandlers,
	}
}
