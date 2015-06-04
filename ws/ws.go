package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleWS(w http.ResponseWriter, r *http.Request) {
	//	io.WriteString(w, "Hello World!")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		defer ws.Close()
		ws.SetReadLimit(maxMessageSize)
		ws.SetReadDeadline(time.Now().Add(pongWait))
		ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				break
			}
			fmt.Println("receive:", string(message))
		}
	}()

	go func() {
		defer ws.Close()
		for {
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.TextMessage, []byte("This Server!")); err != nil {
				fmt.Println(err)
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

}

func main() {
	http.HandleFunc("/ws", HandleWS)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println(err)
	}
}
