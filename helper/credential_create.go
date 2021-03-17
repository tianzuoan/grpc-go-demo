package helper

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

func GetServerCredentialForMethodGrpcServerWithCredential() (credentials.TransportCredentials, error) {
	return credentials.NewServerTLSFromFile("../keys/server.crt", "../keys/server_no_password.key")
}

func GetClientCredentialForMethodGrpcServerWithCredential() (credentials.TransportCredentials, error) {
	return credentials.NewClientTLSFromFile("../keys/server.crt", "yncmall.cn")
}

func GetServerCredentialForMethodGrpcServerWithHttp2() (credentials.TransportCredentials, error) {
	return credentials.NewServerTLSFromFile("../keys/server.crt", "../keys/server_no_password.key")
}

func GetServerCredentialForMethodGrpcServerWithCACertificate() credentials.TransportCredentials {
	certificate, err := tls.LoadX509KeyPair("../ca_cert/server.pem", "../ca_cert/server.key")
	if err != nil {
		log.Fatal("certificate load failed!err:", err)
	}
	certPool := x509.NewCertPool()
	caPemContentBytes, err := ioutil.ReadFile("../ca_cert/ca.pem")
	certPool.AppendCertsFromPEM(caPemContentBytes)
	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
}

func GetClientCredentialForMethodGrpcServerWithCACertificate() credentials.TransportCredentials {
	certificate, err := tls.LoadX509KeyPair("../ca_cert/client.pem", "../ca_cert/client.key")
	if err != nil {
		log.Fatal("certificate load failed!err:", err)
	}
	certPool := x509.NewCertPool()
	caPemContentBytes, err := ioutil.ReadFile("../ca_cert/ca.pem")
	certPool.AppendCertsFromPEM(caPemContentBytes)
	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
}
