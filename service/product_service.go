package service

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
)

type ProductService struct {
}

func NewProductServiceServer() *ProductService {
	return &ProductService{}
}

func (s *ProductService) GetProductInfo(ctx context.Context, req *GetProductInfoRequest) (*GetProductInfoReply, error) {
	return &GetProductInfoReply{
		Product: &Product{
			Id:    req.ProductId,
			Name:  "扯犊子商品",
			Price: rand.Float32(),
			Stock: rand.Int63(),
		},
	}, nil
}

func (s *ProductService) GetProductStock(ctx context.Context, req *GetProductStockRequest) (*GetProductStockReply, error) {
	var stock int64
	if req.ProductArea == ProductArea_A {
		stock = int64(req.GetProductId()) * 100
	} else if req.ProductArea == ProductArea_B {
		stock = int64(req.GetProductId()) * 200
	} else if req.ProductArea == ProductArea_C {
		stock = int64(req.GetProductId()) * 300
	}
	return &GetProductStockReply{Stock: stock}, nil
}

func (s *ProductService) GetProductStockList(ctx context.Context, req *PageRequest) (*GetProductStockListReply, error) {
	var i int32
	lists := make([]*GetProductStockReply, req.GetPageSize())
	for i = 0; i < req.GetPageSize(); i++ {
		lists[i] = &GetProductStockReply{
			Stock: int64(i+1) * 100,
		}
	}
	return &GetProductStockListReply{List: lists}, nil
}

func (s *ProductService) GetProductList(ctx context.Context, req *PageRequest) (*GetProductListReply, error) {
	var i int32
	lists := make([]*Product, req.GetPageSize())
	for i = 0; i < req.GetPageSize(); i++ {
		lists[i] = &Product{
			Id:    i + 1,
			Name:  fmt.Sprintf("baodelu %s", strconv.Itoa(int(i))),
			Price: rand.Float32(),
			Stock: int64(i+1) * 100,
		}
	}
	return &GetProductListReply{List: lists}, nil
}
