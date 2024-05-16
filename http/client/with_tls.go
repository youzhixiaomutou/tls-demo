package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
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

	// 创建 HTTP 客户端
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig, // 传入 TLS 配置
		},
	}
	// 发送 HTTP 请求
	url := "https://localhost:1443/ping"
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// 请求返回 pong
	fmt.Printf("response body = %s\n", body)
}
