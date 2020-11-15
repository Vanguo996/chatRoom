package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 创建升级对象
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// 检测来自前端的请求，
	CheckOrigin: func(r *http.Request) bool { return true },
}

func wsHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("upgrade the connection...")
	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Print(err)
		return
	}
	reader(conn)
}

// 通过websocket连接对象，读取值，返回值
func reader(conn *websocket.Conn) {
	fmt.Print("read message...")
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		log.Print(err)
		return
	}
	// 打印读取到的值
	fmt.Println(string(p))

	if err := conn.WriteMessage(messageType, p); err != nil {
		log.Print(err)
		return
	}
}

func setRouter() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "simple server")
	})

	http.HandleFunc("/ws", wsHandler)
}

func main() {
	setRouter()
	http.ListenAndServe(":8080", nil)
	fmt.Println("Chat App v0.01")
}
