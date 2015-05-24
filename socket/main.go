package main

import (
	"fmt"
	"net"
	"os"
	//"reflect"
	"io"
)

var (
	port    string = "7777"
	lsnr    net.Listener
	stopKey string = "zhimaguanmen"
)

func main() {
	cmd := "start"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "start":
		start()
	case "stop":
		send(stopKey)
	case "send":
		if len(os.Args) > 2 {
			send(os.Args[2])
		} else {
			send("Hello!")
		}
	default:
		fmt.Println("illegal command")
	}
}

func start() {
	var err error
	lsnr, err = net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		panic("Error listening: " + err.Error())
	}
	defer lsnr.Close()
	fmt.Println("Starting the server")

	for {
		conn, err := lsnr.Accept()
		if err != nil {
			panic("Error accept:" + err.Error())
		}
		fmt.Println("Accepted the Connection :", conn.RemoteAddr())

		msg, _ := readConn(conn)

		if msg == stopKey {
			fmt.Println("Stop the server")
			break
		} else {
			fmt.Println(msg)
		}
	}
}

func readConn(conn net.Conn) (msg string, err error) {
	defer conn.Close()

	buf := make([]byte, 256)
	for {
		n, err := conn.Read(buf)
		if err == nil {
			msg += string(buf[0:n])
		} else if err == io.EOF {
			break
		} else {
			fmt.Println("Error read:", err.Error())
			break
		}
	}
	return
}

func send(msg string) {
	conn, err := net.Dial("tcp", "127.0.0.1:"+port)
	if err != nil {
		panic("Error dial: " + err.Error())
	}
	defer conn.Close()
	_, err = conn.Write([]byte(msg))
	if err != nil {
		panic("Error write: " + err.Error())
	}

	fmt.Println("Success send:" + msg)
}
