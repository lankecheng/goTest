package main

import (
	"fmt"
	"time"
	//"reflect"
	//"os"
	//"os/exec"
)

func main() {
	b := make(chan int)
	fmt.Println("1 ", b)
	go func(bb chan int) {
		time.Sleep(time.Second * 2)
		fmt.Println("go ", bb)
		bb <- 1
	}(b)
	close(b)
	b = make(chan int)
	fmt.Println("2 ", b)
	time.Sleep(time.Second * 5)
	fmt.Println("OK: ")
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

func test(err chan error) {
	time.Sleep(time.Second * 2)
	err <- fmt.Errorf("just have a test!")
}

/**
func c(args ...string) {
	fmt.Println(reflect.TypeOf(args))
}
**/
