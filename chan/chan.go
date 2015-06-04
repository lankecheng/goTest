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
	time.Sleep(time.Millisecond * 100)
	fmt.Println("read first time:", <-c)
	time.Sleep(time.Millisecond * 100)
	fmt.Println("read second time:", <-c)
	time.Sleep(time.Millisecond * 100)
	fmt.Println("read third time:", <-c)
	time.Sleep(time.Millisecond * 100)

	fmt.Println("#############")
	s := make(chan int, 2)
	go func() {
		//time.Sleep(10 * time.Millisecond)
		s <- 1
		s <- 0
		s <- 2
		close(s)
	}()
	//close(s)
	time.Sleep(10 * time.Millisecond)
	for v := range s {
		fmt.Println(v)
	}
}
