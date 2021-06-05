package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

func Authorization(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	methods := strings.Split(info.FullMethod, "/")
	fmt.Println(methods)
	if md, ok := metadata.FromIncomingContext(c); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token11")
	} else {
		var userName string
		fmt.Printf("%v\n", md)
		if d, ok := md["username"]; ok {
			userName = d[0]
		} else {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token22")
		}
		fmt.Println(userName)
	}
	return handler(c, req)
}
