package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tianzuoan/grpc-go-demo/service"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	mux := runtime.NewServeMux()
	//dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCredentialForMethodGrpcServerWithCACertificate())}
	dialOptions := []grpc.DialOption{grpc.WithInsecure()}
	err := service.RegisterProductServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		"localhost:8081",
		dialOptions,
	)
	err = service.RegisterOrderServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		"localhost:8081",
		dialOptions,
	)
	if err != nil {
		log.Fatal("http server start failed!err:", err)
	}
	/*****http-server			start*****/
	httpServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Printf("http server start at: %v", "8080")
	_ = httpServer.ListenAndServe()
	/*****http-server			end*****/

	/* 等价上面的 http-server注释块
	log.Printf("http server start at: %v", "8080")
	_ = http.ListenAndServe(":8080", mux)
	*/
}
