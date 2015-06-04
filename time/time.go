package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	/**
	c := make(chan int)
	go func() {
		time.Sleep(time.Second * 1)
		c <- 1
	}()
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
	t1 := time.Now()
	time.Sleep(time.Millisecond * 4323)
	//time.Sleep(time.Second * 3)
	fmt.Printf("%.fms\n", time.Since(t1).Seconds()*1000)
}

func c(args ...string) {
	fmt.Println(reflect.TypeOf(args))
}
