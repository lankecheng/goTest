package main

import (
	"fmt"
	"time"
)

type TEST_VAL string

func main() {
	c := make(chan int, 0)
	go func() {
		//time.Sleep(time.Second * 2)
		c <- 1
		fmt.Println("input 1, chan size ", len(c))
		c <- 2
		fmt.Println("input 2, chan size ", len(c))
		c <- 3
		fmt.Println("input 3")
	}()
	/**
	durTime := time.Duration(time.Duration(1) * time.Second)
	select {
	case v, ok := <-c:
		fmt.Println(v, "-c-", ok)
	case v, ok := <-time.After(durTime):
		fmt.Println(v, "-timeout-", ok)
	}
	**/
	time.Sleep(time.Second * 1)
	fmt.Println("read first time:", <-c)
	time.Sleep(time.Second * 1)
	fmt.Println("read second time:", <-c)
	time.Sleep(time.Second * 1)
	fmt.Println("read third time:", <-c)
	time.Sleep(time.Second * 1)
}
