package service

import repository "test_shop/pkg/repositry"

type Shop interface {
	GenerateListOrders(listOrders *[]int64) (int, error)
}
type Service struct {
	Shop
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Shop: NewShopService(repos.Shop),
	}
}
