package repository

import (
	"github.com/jmoiron/sqlx"
	"test_shop"
)

type Shop interface {
	GetOrderById(id int64) (*[]test_shop.OutPutOnceOrderInfo, error)
}
type Repository struct {
	Shop
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Shop: NewShopPostgres(db),
	}
}
