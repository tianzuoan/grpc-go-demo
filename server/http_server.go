package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tianzuoan/grpc-go-demo/helper"
	"github.com/tianzuoan/grpc-go-demo/service"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	mux := runtime.NewServeMux()
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCredentialForMethodGrpcServerWithCACertificate())}
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
	httpServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("grpc server start at: %v", "8080")

	_ = httpServer.ListenAndServe()
}
