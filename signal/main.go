package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	//"time"
	//"github.com/lankecheng/stringutil"
)

type TEST_VAL string

func main() {
	//fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	//a()
	//b := TEST_VAL("test_val")
	//fmt.Println(b)

	sig_chan := make(chan os.Signal)
	signal.Notify(sig_chan, os.Interrupt, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		sig := <-sig_chan
		fmt.Println("here it is interrupt", sig)
		os.Exit(1)
	}()
	for {
		panic(3354)
		//time.Sleep(time.Second * 1)
	}
	//a := make(chan int)
	//<-a
	//for {
	//}
	/**
	c := make(chan int, 0)
	go func() {
		time.Sleep(time.Second * 2)
		c <- 1
		fmt.Println("input 1, chan size ", len(c))
		c <- 2
		fmt.Println("input 2, chan size ", len(c))
		c <- 3
		fmt.Println("input 3")
	}()
	durTime := time.Duration(time.Duration(1) * time.Second)
	select {
	case v, ok := <-c:
		fmt.Println(v, "-c-", ok)
	case v, ok := <-time.After(durTime):
		fmt.Println(v, "-timeout-", ok)
	}
	time.Sleep(time.Second * 1)
	fmt.Println("read first time:", <-c)
	time.Sleep(time.Second * 1)
	fmt.Println("read second time:", <-c)
	time.Sleep(time.Second * 1)
	fmt.Println("read third time:", <-c)
	time.Sleep(time.Second * 1)
	**/
}

func a() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
}
