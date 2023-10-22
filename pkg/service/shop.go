package service

import (
	"fmt"
	"sort"
	"test_shop"
	repository "test_shop/pkg/repositry"
)

type ShopService struct {
	repo repository.Shop
}

func NewShopService(repo repository.Shop) *ShopService {
	return &ShopService{
		repo: repo,
	}
}

func keys(myMap *map[string][]*test_shop.OutPutOnceOrderInfo) *[]string {
	keyList := make([]string, len(*myMap))

	i := 0
	for k := range *myMap {
		keyList[i] = k
		i++
	}
	return &keyList
}

func printer(myMap *map[string][]*test_shop.OutPutOnceOrderInfo) {
	pointer := keys(myMap)
	//fmt.Println(*pointer)
	sort.Slice(*pointer, func(i, j int) bool {
		return (*pointer)[i] < (*pointer)[j]
	})
	//fmt.Println(*pointer)
	for i := 0; i < len(*pointer); i++ {
		fmt.Printf("===Стеллаж %s\n", (*pointer)[i])
		for j := 0; j < len((*myMap)[(*pointer)[i]]); j++ {
			fmt.Printf("%s (id=%d)\nзаказ %d, %d шт\n", (*myMap)[(*pointer)[i]][j].NameTech,
				(*myMap)[(*pointer)[i]][j].IdOfTech, (*myMap)[(*pointer)[i]][j].OrderId,
				(*myMap)[(*pointer)[i]][j].NumberOfTech)
			if len((*myMap)[(*pointer)[i]][j].NamesOfOtherWarehouses) > 0 {
				fmt.Print("доп стеллаж: ")
				for k := 0; k < len((*myMap)[(*pointer)[i]][j].NamesOfOtherWarehouses)-1; k++ {
					fmt.Printf("%s,", (*myMap)[(*pointer)[i]][j].NamesOfOtherWarehouses[k])
				}
				fmt.Printf("%s\n", (*myMap)[(*pointer)[i]][j].NamesOfOtherWarehouses[len((*myMap)[(*pointer)[i]][j].NamesOfOtherWarehouses)-1])
			}
			fmt.Print("\n")
		}
	}
}

func (s *ShopService) GenerateListOrders(listOrders *[]int64) (int, error) {
	list := make(map[string][]*test_shop.OutPutOnceOrderInfo)
	//var list map[string][]test_shop.OutPutOnceOrderInfo
	for i := 0; i < len(*listOrders); i++ {
		obj, err := s.repo.GetOrderById((*listOrders)[i])
		for j := 0; j < len((*obj)); j++ {
			name := (*obj)[j].NameOfWarehouse
			list[name] = append(list[name], &((*obj)[j]))
		}
		if err != nil {
			fmt.Println(err)
		}
	}
	printer(&list)
	return 0, nil
}
