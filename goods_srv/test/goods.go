package main

import (
	"context"
	"fmt"
	"mxshop_srvs/goods_srv/proto"
)

func TestGetGoodsList() {
	rsp, err := brandClient.GoodsList(context.Background(), &proto.GoodsFilterRequest{
		TopCategory: 135200,
		PriceMin: 60,
		// KeyWords: "三都港",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, good := range rsp.Data {
		fmt.Println(good.Name, good.ShopPrice)
	}
}