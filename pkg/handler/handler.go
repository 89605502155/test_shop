package handler

import (
	"fmt"
	"test_shop/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitConsole(lenth int) (int, error) {
	orders := make([]int64, lenth)
	for i := 0; i < len(orders)-1; i++ {
		fmt.Scanf("%d,", &orders[i])
	}
	fmt.Scanf("%d", &orders[len(orders)-1])
	resp, err := h.services.Shop.GenerateListOrders(&orders)
	return resp, err
}
