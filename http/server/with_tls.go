package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ping-pong handler
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	var (
		certFile = "cert/cert.pem"
		keyFile  = "cert/key.pem"
	)
	// HTTP with TLS on 1443
	fmt.Println("HTTP server with tls is running on :1443.")
	err := http.ListenAndServeTLS(":1443", certFile, keyFile, nil)
	if err != nil {
		panic(err)
	}
}
