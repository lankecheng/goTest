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
	"reflect"
	"strings"
)

func main() {
	b := "abc123"
	fmt.Println(b[0], reflect.TypeOf(b[0]))

	bb := string(b[0])
	fmt.Println(bb, reflect.TypeOf(bb))


	bary := []rune(b)
	chg(bary)

	//	bary[0] = 65

	//	b[0] = b[0] - 32
	fmt.Println(reflect.TypeOf(bary))
	fmt.Println(string(bary))

	fmt.Println(strings.Split(b, "_"))
	fmt.Println("-------")
	fmt.Println("-------")
	birth, _ := time.ParseInLocation("2006-01-02", "2015-06-07", time.Local)
	fmt.Println(birth)
	fmt.Println(birth.Unix())

	fmt.Println(time.Unix(birth.Unix(), 0).UTC())

	fmt.Println("-------")
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	fmt.Println("-------")
	fmt.Println("0123"[0:2])
	fmt.Println("-------")

	var s interface{}
	s = "abc"
	switch t := s.(type) {
	case *string:
		fmt.Println(t)
	case string:
		fmt.Println(t)
	}

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

func chg(r []rune) {
	r[0] = 100
}

/**
func c(args ...string) {
	fmt.Println(reflect.TypeOf(args))
}
**/
