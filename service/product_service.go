package service

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

type CustomResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Detail  string      `json:"detail"`
	Data    interface{} `json:"data"`
}

// GetProductListReply实现自定义grpc-gateway http接口数据返回格式，只需实现
//type responseBody interface {
//   XXX_ResponseBody() interface{}
//}接口
func (x *GetProductListReply) XXX_ResponseBody() interface{} {
	return &CustomResponse{Code: 0, Message: "success", Detail: "", Data: x}
}

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

//服务端流模式
func (s *ProductService) GetProductStockListByServerStream(req *PageRequest, stream ProductService_GetProductStockListByServerStreamServer) error {
	var i, batchSize int32

	lists := make([]*GetProductStockReply, batchSize)
	var count int32
	for i = 1; i <= req.PageSize; i++ {
		lists = append(lists, &GetProductStockReply{
			Stock: int64(i) * 100,
		})
		if i%2 == 0 {
			count++
			err := stream.Send(&GetProductStockListReply{List: lists})
			fmt.Printf("【服务端流模式】服务端开始分批次给客户端返回数据，第 %d 回合，返回的的数据如下：%v \n\n", count, lists)
			if err != nil {
				fmt.Println("stream send failed! err:", err)
				return err
			}
			time.Sleep(time.Millisecond * 500)
			lists = nil
		}
	}
	if lists != nil {
		err := stream.Send(&GetProductStockListReply{List: lists})
		count++
		fmt.Printf("【服务端流模式】服务端开始分批次(最后一次)给客户端返回数据，第 %d 回合，返回的的数据如下：%v \n\n", count, lists)
		if err != nil {
			fmt.Println("stream send failed! err:", err)
			return err
		}
	}
	return nil
}

//客户端流模式
func (s *ProductService) GetProductStockListByClientStream(stream ProductService_GetProductStockListByClientStreamServer) error {
	lists := make([]*GetProductStockReply, 0)
	var count int32
	for {
		count++
		request, err := stream.Recv()
		fmt.Printf("【客户端流模式】服务端开始分批次接收客户端的数据，第 %d 回合，接收到的数据如下：%v \n\n", count, request.GetProductIds())
		if err == io.EOF { //代表接收完了
			fmt.Println("【客户端流模式】数据接收完了，一次性把所有处理好的数据返回给客户端！")
			return stream.SendAndClose(&GetProductStockListReply{List: lists})
		}

		if err != nil {
			return err
		}

		for _, productId := range request.GetProductIds() {
			lists = append(lists, &GetProductStockReply{
				ProductId: productId,
				Stock:     int64(productId*100 + 1),
			})
		}
		time.Sleep(time.Millisecond * 500)
	}
}

//双向流模式
func (s *ProductService) GetProductStockListByTwinsStream(stream ProductService_GetProductStockListByTwinsStreamServer) error {
	lists := make([]*GetProductStockReply, 0)
	var count int32
	for {
		count++
		request, err := stream.Recv()
		fmt.Printf("【双向流模式】服务端开始分批次接收客户端的数据，第 %d 回合，接收到的数据如下：%v \n\n", count, request.GetProductIds())
		if err == io.EOF { //代表接收完了
			fmt.Println("【双向流模式】没有数据接收了,终结！")
			return nil
		}

		if err != nil {
			return err
		}

		for _, productId := range request.GetProductIds() {
			lists = append(lists, &GetProductStockReply{
				ProductId: productId,
				Stock:     int64(productId*100 + 1),
			})
		}
		time.Sleep(time.Millisecond * 500)
		err = stream.Send(&GetProductStockListReply{List: lists})
		fmt.Printf("【双向流模式】服务端开始分批次把第 %d 回合接收的数据处理好后发送给客户端，第 %d 回合，发送给客户端的数据如下：%v \n\n", count, count, lists)
		if err != nil {
			return err
		}
		lists = lists[0:0]
	}
}
