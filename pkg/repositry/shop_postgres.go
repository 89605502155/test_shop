package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"test_shop"
)

type ShopPostgres struct {
	db *sqlx.DB
}

func NewShopPostgres(db *sqlx.DB) *ShopPostgres {
	return &ShopPostgres{db: db}
}
func (r *ShopPostgres) getOrderListById(id *int64) (*test_shop.Order, error) {
	order := test_shop.Order{}
	queryKindsOfTech := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", orderTable)
	err := r.db.Get(&order, queryKindsOfTech, *id)
	return &order, err
}
func (r *ShopPostgres) getAdditionalShelf(idWarehouses *pq.Int64Array, res *test_shop.OutPutOnceOrderInfo) {
	if len(*idWarehouses) > 0 {
		(*res).NamesOfOtherWarehouses = make([]string, 0)
		//var rrr pq.StringArray
		var dopName string
		for j := 0; j < len(*idWarehouses); j++ {
			queryKinndsOfTech_ := fmt.Sprintf("SELECT name FROM %s WHERE id=$1",
				warehouseTable)
			if (*idWarehouses)[j] != 0 {
				_ = r.db.Get(&dopName, queryKinndsOfTech_, (*idWarehouses)[j])
				(*res).NamesOfOtherWarehouses = append((*res).NamesOfOtherWarehouses, dopName)
			}

		}
	}
}
func (r *ShopPostgres) cicleOfGenerateListShelfs(res *test_shop.OutPutOnceOrderInfo,
	product *test_shop.Product, resoult *[]test_shop.OutPutOnceOrderInfo,
) {
	queryKinndsOfTech_1 := fmt.Sprintf("SELECT name FROM %s WHERE id=$1", warehouseTable)
	_ = r.db.Get(&(*res).NameOfWarehouse, queryKinndsOfTech_1, (*product).MainWarehouseId)

	r.getAdditionalShelf(&(*product).IdWarehouses, res)
	(*resoult) = append((*resoult), (*res))
}
func (r *ShopPostgres) GetOrderById(id int64) (*[]test_shop.OutPutOnceOrderInfo, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	order, err := r.getOrderListById(&id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//numberOfKindOfTech:=len((*order).KindOfTech)
	resoult := make([]test_shop.OutPutOnceOrderInfo, 0)
	for i := 0; i < len((*order).KindOfTech); i++ {

		res_ := test_shop.OutPutOnceOrderInfo{}
		res_.OrderId = id
		res_.IdOfTech = (*order).KindOfTech[i]

		product := test_shop.Product{}
		queryKindsOfTech := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", productTable)
		err = r.db.Get(&product, queryKindsOfTech, res_.IdOfTech)

		res_.NameTech = product.Name
		for {
			res := res_
			if (*order).NumberOfTech[i] == 0 {
				break
			}
			if product.NumberOfProductsOnMainWarehouse > (*order).NumberOfTech[i] {
				res.NumberOfTech = (*order).NumberOfTech[i]
				product.NumberOfProductsOnMainWarehouse -= (*order).NumberOfTech[i]
				(*order).NumberOfTech[i] = 0
				r.cicleOfGenerateListShelfs(&res, &product, &resoult)
				updateQuery := fmt.Sprintf("UPDATE %s SET number_of_products_on_main_warehouse=$1  WHERE id=%d",
					productTable, product.Id)
				_, err = r.db.Exec(updateQuery, product.NumberOfProductsOnMainWarehouse)

			} else {
				if product.NumberOfProductsOnMainWarehouse > 0 {
					res.NumberOfTech = product.NumberOfProductsOnMainWarehouse
					(*order).NumberOfTech[i] -= product.NumberOfProductsOnMainWarehouse
					r.cicleOfGenerateListShelfs(&res, &product, &resoult)
					if len(product.IdWarehouses) > 1 {
						updateQuery := fmt.Sprintf("UPDATE %s SET number_of_products_on_main_warehouse=number_products_on_warehouse[1], main_warehouse_id=id_warehouses[1], number_products_on_warehouse=number_products_on_warehouse[2:],id_warehouses=id_warehouses[2:] WHERE id=$1",
							productTable)
						_, err = r.db.Exec(updateQuery, product.Id)

						product.MainWarehouseId, product.NumberOfProductsOnMainWarehouse = product.IdWarehouses[0], product.NumberProductsOnWarehouse[0]
						product.IdWarehouses = append(product.IdWarehouses[1:])
						product.NumberProductsOnWarehouse = append(product.NumberProductsOnWarehouse[1:])
					} else if len(product.IdWarehouses) == 1 {
						updateQuery := fmt.Sprintf("UPDATE %s SET number_of_products_on_main_warehouse=number_products_on_warehouse[1], main_warehouse_id=id_warehouses[1], number_products_on_warehouse=null,id_warehouses=null WHERE id=$1",
							productTable)
						_, err = r.db.Exec(updateQuery, product.Id)

						if err != nil {
							fmt.Println("dddd")
						}
						product.MainWarehouseId, product.NumberOfProductsOnMainWarehouse = product.IdWarehouses[0], product.NumberProductsOnWarehouse[0]
						product.IdWarehouses = nil
						product.NumberProductsOnWarehouse = nil
					} else {
						if product.NumberOfProductsOnMainWarehouse < (*order).NumberOfTech[i] {
							fmt.Printf("Мы смогли выполнить заказ %d частично. Нам не хватает %d %s.", (*order).Id, (*order).NumberOfTech[i], product.Name)
						}
						(*order).NumberOfTech[i] = 0
						updateQuery := fmt.Sprintf("UPDATE %s SET number_of_products_on_main_warehouse=0 WHERE id=$1", productTable)
						_, err = r.db.Exec(updateQuery, product.Id)

					}
				} else if product.NumberOfProductsOnMainWarehouse == 0 {
					fmt.Printf("Мы не можем выполнить заказ %d. Техника %s закончилась.", (*order).Id, product.Name)
					(*order).NumberOfTech[i] = 0
				} else {
					(*order).NumberOfTech[i] = 0
					tx.Rollback()
					return nil, err
				}
			}
		}
	}
	return &resoult, tx.Commit()
}
