package main

import (
	"fmt"
	//"sync"
	//"reflect"
	"time"
	//"os"
	//"os/exec"
	//"sort"

	//"bytes"
	//"io/ioutil"
	//"io"
	//"net/http"
)

func main() {
	//url := "http://mvvideo4.meitudata.com/539817f2c44405901.mp4"
	url := "http://mvvideo2.meitudata.com/556562923dab95528.mp4"
	fmt.Println(len(url))
	fmt.Println(len([]byte(url)))

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
