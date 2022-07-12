package handler

import (
	"context"

	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"mxshop_srvs/goods_srv/proto"
)

func (s *GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	brandListResponse := proto.BrandListResponse{}

	var brands []model.Brands
	result := global.DB.Scopes(model.Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}

	var total int64
	global.DB.Model(&model.Brands{}).Count(&total)
	brandListResponse.Total = int32(total)

	var brandResponse []*proto.BrandInfoResponse
	for _, brand := range brands {
		brandResponse = append(brandResponse, &proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}
	brandListResponse.Data = brandResponse
	return &brandListResponse, nil
}

// CreateBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*BrandInfoResponse, error)
// DeleteBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
// UpdateBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
