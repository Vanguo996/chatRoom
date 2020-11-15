package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func wsHandler(rw http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrader(rw, r)

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
