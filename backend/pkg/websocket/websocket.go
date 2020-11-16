package websocket

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

func Upgrader(rw http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	fmt.Println("upgrade the connection...")
	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Println(err)
		return conn, err
	}
	return conn, nil
}

// // Reader: 通过websocket连接对象，读取值，返回值
// func Reader(conn *websocket.Conn) {
// 	fmt.Print("read message...")
// 	messageType, p, err := conn.ReadMessage()
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	// 打印读取到的值
// 	fmt.Println("server read message: ", string(p))

// 	if err := conn.WriteMessage(messageType, p); err != nil {
// 		log.Print(err)
// 		return
// 	}
// }

// func Writer(conn *websocket.Conn) {
// 	// messageconn.NextWriter()

// 	for {
// 		fmt.Println("sending message...")
// 		messageType, r, err := conn.NextReader()
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		w, err := conn.NextWriter(messageType)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		if _, err := io.Copy(w, r); err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		if err := w.Close(); err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 	}

// }
