package main

import (
	"mxshop_srvs/goods_srv/proto"

	"google.golang.org/grpc"
)

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	brandClient = proto.NewGoodsClient(conn)
}

func main() {
	Init()
	// TestGetBrandList()
	// TestGetCategoryList()
	TestGetSubCategoryList()
	conn.Close()
}