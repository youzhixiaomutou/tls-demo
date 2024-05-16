package main

import (
	"context"
	"fmt"

	v1 "tls-demo/grpc/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	var (
		certFile = "cert/cert.pem"
	)
	// 加载 TLS 证书
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		panic(err)
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.NewClient("localhost:1443", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建 GRPC 客户端
	c := v1.NewPingPongServerClient(conn)
	// 向服务端发送 ping
	reply, err := c.PingPong(context.Background(), &v1.Request{Message: "ping"})
	if err != nil {
		panic(err)
	}
	// 客户端收到 pong
	fmt.Println(reply)
}
