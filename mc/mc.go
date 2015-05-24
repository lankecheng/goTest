package main

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	//"os"
	"net"
)

var (
	url string = "192.168.41.45:22202"
)

func main() {
	//checkMC()
	mc := memcache.New(url)
	//mc := memcache.New("192.168.41.13:11001")
	item, err := mc.Get("q1")
	if err == nil {
		fmt.Println(string(item.Value))
	} else {
		fmt.Println(err.Error())
	}
}

func checkMC() {
	conn, err := net.Dial("tcp", url)
	if err == nil {
		conn.Close()
	} else {
		panic(err.Error())
	}
}
