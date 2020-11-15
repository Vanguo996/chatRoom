package main

import (
	"fmt"
	"net/http"

	"chatRoom/backend/pkg/websocket"
)

func wsHandler(rw http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrader(rw, r)
	if err != nil {
		fmt.Fprintf(rw, "%+V\n", err)
	}

	go websocket.Writer(conn)
	websocket.Reader(conn)

}

func setRouter() {

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(rw, "simple server")
	// })

	http.HandleFunc("/ws", wsHandler)
}

func main() {
	setRouter()
	http.ListenAndServe(":8080", nil)
	fmt.Println("Chat App v0.01")
}
