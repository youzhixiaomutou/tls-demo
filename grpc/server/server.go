package main

import (
	"context"
	"fmt"
	"net"

	v1 "tls-demo/grpc/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	var (
		certFile = "cert/cert.pem"
		keyFile  = "cert/key.pem"
	)
	// 加载 TLS 证书
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		panic(err)
	}

	// 监听 1443 端口
	lis, err := net.Listen("tcp", ":1443")
	if err != nil {
		panic(err)
	}
	// 创建 GRPC 服务端
	s := grpc.NewServer(grpc.Creds(creds))
	v1.RegisterPingPongServerServer(s, &pingPongServer{})

	fmt.Println("GRPC server with tls is running on :1443.")
	_ = s.Serve(lis)
}

var _ v1.PingPongServerServer = (*pingPongServer)(nil)

type pingPongServer struct {
	v1.UnimplementedPingPongServerServer
}

func (s *pingPongServer) PingPong(ctx context.Context, req *v1.Request) (*v1.Response, error) {
	fmt.Println(req.Message)
	// 服务端返回 pong
	return &v1.Response{Message: "pong"}, nil
}
