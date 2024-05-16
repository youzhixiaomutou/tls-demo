package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	// ping-pong websocket handler
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer ws.Close()
		mt, msg, _ := ws.ReadMessage()
		log.Printf("Received message: %s", msg)
		_ = ws.WriteMessage(mt, []byte("pong"))
	})

	var (
		certFile = "cert/cert.pem"
		keyFile  = "cert/key.pem"
	)
	// WebSocket with TLS on 1443
	fmt.Println("WebSocket server with tls is running on :1443.")
	err := http.ListenAndServeTLS(":1443", certFile, keyFile, nil)
	if err != nil {
		panic(err)
	}
}
