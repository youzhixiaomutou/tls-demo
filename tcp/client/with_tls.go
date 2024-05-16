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

	// 建立 TCP 连接
	conn, err := net.Dial("tcp", ":1443")
	if err != nil {
		panic(err)
	}
	// 使用 TLS 客户端来包装 TCP 连接使其支持 TLS
	tlsConn := tls.Client(conn, tlsConfig)
	defer tlsConn.Close()
	// 执行 TLS 握手
	err = tlsConn.Handshake()
	if err != nil {
		panic(err)
	}

	// 发送 ping 给服务端
	_, _ = tlsConn.Write([]byte("ping"))
	// 服务端返回 pong
	var buff [64]byte
	n, _ := tlsConn.Read(buff[:])
	fmt.Printf("Received from server: %s\n", buff[:n])

}
