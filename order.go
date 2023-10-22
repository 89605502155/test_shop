package test_shop

import "github.com/lib/pq"

type Warehouse struct {
	Id   int64  `json:"-" db:"id"`
	Name string `json:"name" db:"name"`
}
type Product struct {
	Id                              int64         `json:"-" db:"id"`
	Name                            string        `json:"name" db:"name"`
	MainWarehouseId                 int64         `json:"main_warehouse_id" db:"main_warehouse_id"`
	NumberOfProductsOnMainWarehouse int64         `json:"number_of_products_on_main_warehouse" db:"number_of_products_on_main_warehouse"`
	IdWarehouses                    pq.Int64Array `json:"id_warehouses" db:"id_warehouses"`
	NumberProductsOnWarehouse       pq.Int64Array `json:"number_products_on_warehouse" db:"number_products_on_warehouse"`
}
type Order struct {
	Id           int64         `json:"-" db:"id"`
	KindOfTech   pq.Int64Array `json:"kind_of_tech" db:"kind_of_tech"`
	NumberOfTech pq.Int64Array `json:"number_of_tech" db:"number_of_tech"`
}
type InputOrdersSlice struct {
	Id []int `json:"id" db:"id"`
}
type OutPutOnceOrderInfo struct {
	NameTech               string
	OrderId                int64
	IdOfTech               int64
	NumberOfTech           int64
	NameOfWarehouse        string
	NamesOfOtherWarehouses []string
}
type OutputProductCard struct {
	NameTech       string `json:"name_tech" db:"name"`
	IdOfKindOfTech int    `json:"id_of_kind_of_tech" db:"id"`
	NumberOfTech   int    `json:"number_of_tech"`
	NumberOfOrder  int    `json:"number_of_order" db:"id"`
}
type ShelfProductSlice struct {
	ShelfName string              `json:"shelf_name" db:"name"`
	Products  []OutputProductCard `json:"products"`
}
