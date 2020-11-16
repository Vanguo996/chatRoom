package main

import (
	"fmt"
	"net/http"

	"chatRoom/backend/pkg/websocket"
)

func wsHandler(rw http.ResponseWriter, r *http.Request, pool *websocket.Pool) {
	conn, err := websocket.Upgrader(rw, r)
	if err != nil {
		fmt.Fprintf(rw, "%+V\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()

}

func setRouter() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		wsHandler(rw, r, pool)
	})
}

func main() {
	setRouter()
	http.ListenAndServe(":8080", nil)
	fmt.Println("Chat App v0.01")
}
