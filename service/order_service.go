package service

import (
	"context"
	"net/http"
)

type OrderService struct {
}

func NewOrderServiceServer() *OrderService {
	return &OrderService{}
}

func (s *OrderService) NewOrder(ctx context.Context, req *NewOrderRequest) (*NewOrderReply, error) {
	return &NewOrderReply{
		Status:  http.StatusOK,
		Message: "order created success!",
		Data: &Order{
			OrderId:      req.Order.OrderId,
			OrderNo:      req.Order.OrderNo,
			UserId:       req.Order.UserId,
			OrderPayment: req.Order.OrderPayment,
			OrderTime:    req.Order.OrderTime,
		},
	}, nil
}
