package main

import (
	"github.com/tianzuoan/grpc-go-demo/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const PORT = ":10051"

func main() {
	//监听端口
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		return
	}
	//创建一个grpc 服务器
	grpcServer := grpc.NewServer()
	//注册事件
	service.RegisterTinaServer(grpcServer, service.NewTinaService())

	//gRPC服务反射,方便grpcurl调试
	reflection.Register(grpcServer)

	log.Printf("grpc server start at: %v", PORT)
	//处理链接
	_ = grpcServer.Serve(lis)
}
