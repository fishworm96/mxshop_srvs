package main

import (
	"context"
	"fmt"
	"mxshop_srvs/goods_srv/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

func TestGetCategoryList() {
	rsp, err := brandClient.GetAllCategorysList(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.JsonData)
	// for _, category := range rsp.Data {
	// 	fmt.Println(category.Name)
	// }
}

func TestGetSubCategoryList() {
	rsp, err := brandClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: 135200,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.SubCategorys)
}
