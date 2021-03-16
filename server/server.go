package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/tianzuoan/grpc-go-demo/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

const PORT = ":8080"

func main() {
	//grpcServerWithHttp2()
	grpcServerWithCACertificate()
}

func grpcServerInsecure() {
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

func grpcServerWithCredential() {
	//监听端口
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		return
	}
	//创建https通行证
	transportCredentials, err := credentials.NewServerTLSFromFile("../keys/server.crt", "../keys/server_no_password.key")
	if err != nil {
		log.Fatal("certificate created failed!")
	}
	//创建一个grpc 服务器（带证书）
	grpcServer := grpc.NewServer(grpc.Creds(transportCredentials))
	//注册事件
	service.RegisterTinaServer(grpcServer, service.NewTinaService())

	//gRPC服务反射,方便grpcurl调试
	reflection.Register(grpcServer)

	log.Printf("grpc server start at: %v", PORT)
	//处理链接
	_ = grpcServer.Serve(lis)
}

func grpcServerWithHttp2() {
	//创建https通行证
	transportCredentials, err := credentials.NewServerTLSFromFile("../keys/server.crt", "../keys/server_no_password.key")
	if err != nil {
		log.Fatal("certificate created failed!")
	}
	//创建一个grpc 服务器（带证书）
	grpcServer := grpc.NewServer(grpc.Creds(transportCredentials))
	//注册事件
	service.RegisterTinaServer(grpcServer, service.NewTinaService())

	//gRPC服务反射,方便grpcurl调试
	reflection.Register(grpcServer)

	log.Printf("grpc server start at: %v", PORT)
	//处理链接
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Proto) //HTTP/2.0
		fmt.Println(request.RequestURI)
		fmt.Println(request)
		grpcServer.ServeHTTP(writer, request)
	})
	//创建http服务
	httpServer := http.Server{
		Addr:    PORT,
		Handler: mux,
	}
	//http服务启动监听并带证书认证
	_ = httpServer.ListenAndServeTLS("../keys/server.crt", "../keys/server_no_password.key")
}

func grpcServerWithCACertificate() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		return
	}

	certificate, err := tls.LoadX509KeyPair("../ca_cert/server.pem", "../ca_cert/server.key")
	if err != nil {
		log.Fatal("certificate load failed!err:", err)
	}
	certPool := x509.NewCertPool()
	caPemContentBytes, err := ioutil.ReadFile("../ca_cert/ca.pem")
	certPool.AppendCertsFromPEM(caPemContentBytes)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	//创建一个grpc 服务器（带证书）
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	//注册事件
	service.RegisterTinaServer(grpcServer, service.NewTinaService())

	//gRPC服务反射,方便grpcurl调试
	reflection.Register(grpcServer)

	log.Printf("grpc server start at: %v", PORT)
	//处理链接
	_ = grpcServer.Serve(lis)
}
