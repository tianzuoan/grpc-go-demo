package main

import (
	"context"
	Greeter "github.com/tianzuoan/grpc-go-demo/proto"
	"google.golang.org/grpc"
	"log"
)

const (
	ADDRESS = "localhost:50051"
)

func main() {
	//通过grpc 库 建立一个连接
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	//通过刚刚的连接 生成一个client对象。
	c := Greeter.NewGreeterClient(conn)
	//直接通过 client对象 调用 服务端的函数
	r, err := c.SayHello(context.Background(), &Greeter.HelloRequest{Name: "namas"})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
	langReply, err := c.LearnForeignLanguage(context.Background(), &Greeter.LanguageRequest{Kind: "English"})
	if err != nil {
		log.Fatal("could not learn language,reason is :%v", err)
	}
	log.Printf("learn language,the score is:%v", langReply.Score)
}
