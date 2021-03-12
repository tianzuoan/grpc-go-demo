package main

import (
	"context"
	"github.com/tianzuoan/grpc-go-demo/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const (
	ADDRESS = ":8080"
)

func main() {
	grpcServerWithCredential()
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
	transportCredentials, err := credentials.NewClientTLSFromFile("../keys/server.crt", "yncmall.cn")
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
