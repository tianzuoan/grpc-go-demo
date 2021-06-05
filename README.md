# grpc-go-demo
golang grpc example with grpc server reflection

how to run this demo
1.  cd proto/ && protoc --go_out=plugins=grpc:../ --grpc-gateway_out=logtostderr=true:../ *.proto && cd ../
2.  go run server/server.go
3.  go run client/client.go