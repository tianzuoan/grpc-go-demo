package main

import (
	"context"
	Greeter "github.com/tianzuoan/grpc-go-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const PORT = ":50051"

type server struct{}

func (s *server) SayHello(ctx context.Context, in *Greeter.HelloRequest) (*Greeter.HelloReply, error) {
	return &Greeter.HelloReply{Message: in.GetName() + "hello"}, nil
}

func (s *server) LearnForeignLanguage(ctx context.Context, in *Greeter.LanguageRequest) (*Greeter.LanguageReply, error) {
	return &Greeter.LanguageReply{Score: 150}, nil
}

func main() {
	//监听端口
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		return
	}
	//创建一个grpc 服务器
	s := grpc.NewServer()
	//注册事件
	Greeter.RegisterGreeterServer(s, &server{})

	//gRPC服务反射
	reflection.Register(s)

	log.Printf("grpc server start at: %v", PORT)
	//处理链接
	_ = s.Serve(lis)
}
