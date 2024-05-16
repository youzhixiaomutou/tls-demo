package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
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

	// 传入 TLS 配置
	websocket.DefaultDialer.TLSClientConfig = tlsConfig
	wssURL := "wss://localhost:1443/ping"
	ws, _, err := websocket.DefaultDialer.Dial(wssURL, http.Header{})
	if err != nil {
		panic(err)
	}
	defer ws.Close()

	//  发送 ping 给服务端
	_ = ws.WriteMessage(websocket.TextMessage, []byte("ping"))
	// 服务端返回 pong
	_, p, _ := ws.ReadMessage()
	fmt.Printf("Received message: %s\n", p)
}
