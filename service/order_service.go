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

func (s *OrderService) NewOrder(ctx context.Context, req *Order) (*NewOrderReply, error) {
	return &NewOrderReply{
		Status:  http.StatusOK,
		Message: "order created success!",
		Data: &Order{
			OrderId:      req.OrderId,
			OrderNo:      req.OrderNo,
			UserId:       req.UserId,
			OrderPayment: req.OrderPayment,
			OrderTime:    req.OrderTime,
		},
	}, nil
}
