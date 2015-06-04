package main

import (
	"fmt"
	//"sync"
	//"reflect"
	"time"
	//"os"
	//"os/exec"
	//"sort"

	"net/http"

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

func main() {
	d := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
		ReadBufferSize:   512,
	}

	hMap := make(map[string][]string)
	hMap["Upgrade"] = []string{"websocket"}
	hMap["Connection"] = []string{"Upgrade"}
	hMap["Sec-WebSocket-Key"] = []string{"U00QUfV1CRfIIU0NkcUCnA=="}
	hMap["Sec-WebSocket-Version"] = []string{"13"}
	hMap["Sec-WebSocket-Extensions"] = []string{"x-webkit-deflate-frame"}
	
	h := http.Header(hMap)

	ws, _, err := d.Dial("ws://127.0.0.1:9999/ws", h)
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}

	shutdown := make(chan struct{})

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
		shutdown <- struct{}{}
	}()

	go func() {
		defer ws.Close()
		for {
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.TextMessage, []byte("This Client!")); err != nil {
				fmt.Println(err)
				return
			}
			time.Sleep(1 * time.Second)
		}
		shutdown <- struct{}{}
	}()

	<-shutdown
	fmt.Println("shutdown")
	/**
	file, _ := os.Getwd()
	fmt.Println("current path:", file)
	file, _ = exec.LookPath(os.Args[0])
	fmt.Println("exec path:", file)
	**/
	//fmt.Printf("%v", reflect.TypeOf(v))
	//c <- "a"
	/**
	select {
	case <-c:
	case <-time.After(time.Second * 3):
		fmt.Println("time out")
	}
	**/
	/**
	args := []string{"a", "b"}
	c(args[0:])
	**/
}

func test2(arg string) (filePath []string) {
	fmt.Println(filePath == nil)
	return
}

func test(err chan error) {
	time.Sleep(time.Second * 2)
	err <- fmt.Errorf("just have a test!")
}

/**
func c(args ...string) {
	fmt.Println(reflect.TypeOf(args))
}
**/
