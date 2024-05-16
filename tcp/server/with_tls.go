package main

import (
	"crypto/tls"
	"fmt"
	"net"
)

func main() {
	var (
		certFile = "cert/cert.pem"
		keyFile  = "cert/key.pem"
	)
	// 加载证书和私钥
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		panic(err)
	}
	// 创建 TLS 配置
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert}, // 添加客户端证书
		InsecureSkipVerify: true,                    // 跳过安全性校验，自签名证书必须跳过校验才能使用，生产环境如果使用权威证书则考虑关闭这个选项
	}

	// 创建 TCP listener 监听 1443 端口
	fmt.Println("TCP server with tls is running on :1443.")
	listener, err := net.Listen("tcp", ":1443")
	if err != nil {
		panic(err)
	}
	// 使用 TLS listener 来包装 TCP listener，使其支持 TLS
	tlsListener := tls.NewListener(listener, tlsConfig)
	// 简单处理 TCP 连接请求
	for {
		conn, err := tlsListener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// 发送 pong 给客户端
	_, _ = conn.Write([]byte("pong"))
}
