package main

import (
	"context"
	"fmt"
	"github.com/tianzuoan/grpc-go-demo/helper"
	"github.com/tianzuoan/grpc-go-demo/service"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

const (
	ADDRESS = ":8081"
)

func main() {
	//grpcServerWithCredential()
	grpcServerWithCACertificate()
}

func grpcServerInsecure() {
	//通过grpc 库 建立一个连接
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	//通过刚刚的连接 生成一个client对象。
	c := service.NewTinaClient(conn)
	//直接通过 client对象 调用 服务端的函数
	r, err := c.SayHello(context.Background(), &service.HelloRequest{Name: "马德福"})
	if err != nil {
		log.Fatal("tina service could not say hello:", err)
	}
	log.Printf("tina say hello: %s", r.Message)
	langReply, err := c.LearnForeignLanguage(context.Background(), &service.LanguageRequest{Kind: "chinese"})
	if err != nil {
		log.Fatal("could not learn language,reason is :%v", err)
	}
	log.Printf("learn language,the score is:%v", langReply.Score)
}

func grpcServerWithCredential() {
	//创建凭证
	transportCredentials, err := helper.GetClientCredentialForMethodGrpcServerWithCredential()
	if err != nil {
		log.Fatal("certificated authorized failed!", err)
	}
	//通过grpc 库 建立一个连接（带凭证）
	conn, err := grpc.Dial(ADDRESS, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		return
	}
	defer conn.Close()

	//通过刚刚的连接 生成一个client对象。
	c := service.NewTinaClient(conn)
	//直接通过 client对象 调用 服务端的函数
	r, err := c.SayHello(context.Background(), &service.HelloRequest{Name: "马德福"})
	if err != nil {
		log.Fatal("tina service could not say hello:", err)
	}
	log.Printf("tina say hello: %s", r.Message)
	langReply, err := c.LearnForeignLanguage(context.Background(), &service.LanguageRequest{Kind: "chinese"})
	if err != nil {
		log.Fatal("could not learn language,reason is :%v", err)
	}
	log.Printf("learn language,the score is:%v", langReply.Score)
}

func grpcServerWithCACertificate() {
	//创建凭证
	creds := helper.GetClientCredentialForMethodGrpcServerWithCACertificate()

	//通过grpc 库 建立一个连接（带凭证）
	conn, err := grpc.Dial(ADDRESS, grpc.WithTransportCredentials(creds))
	if err != nil {
		return
	}
	defer conn.Close()

	//通过刚刚的连接 生成一个client对象。
	c := service.NewTinaClient(conn)
	//直接通过 client对象 调用 服务端的函数
	r, err := c.SayHello(context.Background(), &service.HelloRequest{Name: "马德福"})
	if err != nil {
		log.Fatal("tina service could not say hello:", err)
	}
	log.Printf("tina say hello: %s", r.Message)
	langReply, err := c.LearnForeignLanguage(context.Background(), &service.LanguageRequest{Kind: "chinese"})
	if err != nil {
		log.Fatal("could not learn language,reason is :%v", err)
	}
	log.Printf("learn language,the score is:%v", langReply.Score)
	err = getProductStockListByServerStream(conn)
	if err != nil {
		log.Fatal("get product stock list by server stream failed! err:", err)
	}
	err = getProductStockListByClientStream(conn)
	if err != nil {
		log.Fatal("get product stock list by client stream failed! err:", err)
	}
	err = getProductStockListByTwinsStream(conn)
	if err != nil {
		log.Fatal("get product stock list by twins stream failed! err:", err)
	}
}

// 服务端流模式
func getProductStockListByServerStream(conn *grpc.ClientConn) error {
	productClient := service.NewProductServiceClient(conn)
	byServerStreamClient, err := productClient.GetProductStockListByServerStream(context.Background(), &service.PageRequest{PageSize: 5})
	if err != nil {
		return err
	}
	fmt.Println("----------------get product stock list by server stream-----start-------------")

	var count int32
	for {
		count++
		reply, err := byServerStreamClient.Recv()
		if reply != nil {
			fmt.Printf("【服务端流模式】客户端开始分批次接收服务端返回的数据，第 %d 回合，接收的数据如下：%v \n\n", count, reply.List)
		}

		if err == io.EOF {
			fmt.Printf("【服务端流模式】客户端接收数据完毕！\n\n")
			break
		}
	}

	fmt.Println("----------------get product stock list by server stream-----end-------------")
	return nil
}

// 客户端流模式
func getProductStockListByClientStream(conn *grpc.ClientConn) error {
	productClient := service.NewProductServiceClient(conn)
	stream, err := productClient.GetProductStockListByClientStream(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("----------------get product stock list by client stream-----start-------------")
	var count int32
	for j := 0; j < 8; j++ {
		count++
		batchSize := 5
		productIds := make([]int32, batchSize)
		for i := 0; i < batchSize; i++ {
			productIds[i] = int32(j*batchSize + i)
		}
		err := stream.Send(&service.ProductIdsRequest{ProductIds: productIds})
		fmt.Printf("【客户端流模式】客户端开始分批次给服务端发送数据，第 %d 回合，发送的数据如下：%v \n\n", count, productIds)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 300)
	}

	reply, err := stream.CloseAndRecv()
	if reply != nil {
		fmt.Printf("【客户端流模式】客户端一次性从服务端接收数据并关闭流请求，从服务端接收的接收数据如下：%v \n\n", reply.GetList())
	}

	fmt.Println("----------------get product stock list by client stream-----end-------------")
	return nil
}

// 双向流模式
func getProductStockListByTwinsStream(conn *grpc.ClientConn) error {
	productClient := service.NewProductServiceClient(conn)
	stream, err := productClient.GetProductStockListByTwinsStream(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("----------------get product stock list by twins stream-----start-------------")
	var count int32
	for j := 0; j < 8; j++ {
		count++
		batchSize := 5
		productIds := make([]int32, batchSize)
		for i := 0; i < batchSize; i++ {
			productIds[i] = int32(j*batchSize + i)
		}
		err := stream.Send(&service.ProductIdsRequest{ProductIds: productIds})
		fmt.Printf("【双向流模式】客户端开始分批次给服务端发送数据，第 %d 回合，发送的数据如下：%v \n\n", count, productIds)
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 500)
		reply, err := stream.Recv()
		if err != nil {
			return err
		}
		if reply != nil {
			fmt.Printf("【双向流模式】客户端第 %d 回合发送数据后等待服务端该回合数据返回，服务端该回合返回的数据如下：%v \n\n", count, reply.List)
		}
	}

	err = stream.CloseSend()
	fmt.Printf("【双向流模式】客户端结束流请求！ \n\n")
	if err != nil {
		return err
	}

	fmt.Println("----------------get product stock list by twins stream-----end-------------")
	return nil
}
