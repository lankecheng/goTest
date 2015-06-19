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
	"encoding/binary"
	"os"
	"os/exec"
	"github.com/kardianos/osext"
	"path/filepath"
	"path"
	"runtime"
)

func main() {
	var i_a int = 3

	if float64(i_a) > 1.1 {
		fmt.Println(i_a)
	}
	fmt.Println("-------")
	fmt.Println(os.Args[0])
	fmt.Println("-------")
	a, _ := filepath.Abs("./logs")
	fmt.Println("test:", a)
	_, currFile, _, _ := runtime.Caller(0)
	fmt.Println(currFile)
	fmt.Println("-------")
	file, _ := os.Getwd()
	fmt.Printf("current path: %v\n", file)
	file, _ = exec.LookPath(os.Args[0])
	args0, _ := filepath.Abs(file)
	fmt.Printf("args[0]: %v\n", file)
	fmt.Printf("args[0] path: %v\n", args0)
	fmt.Printf("args[0] dir: %v\n", path.Dir(args0))
	logDir, err := filepath.Abs(path.Dir(args0) + "/..")
	if err == nil {
		fmt.Printf("log dir: %v\n", logDir)
	} else {
		fmt.Printf("log dir: %v\n", err)
	}

	fmt.Printf("tmp path: %v\n", os.TempDir())

	filename, _ := osext.Executable()
//	binFolder := path.Dir(filename);
	fmt.Println(filename)

	fmt.Println("-------")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered ", err)
		}
	}()

	fmt.Println([]byte("abc"))
	fmt.Println("-------")
	var test1 int64
	var test2 uint64 = 1
	test1 = int64(test2)
	fmt.Println(test1)
	fmt.Println("-------")
	var i64 int64 = 256
	i64LtlB := make([]byte, 8)
	i64bigB := make([]byte, 8)
	binary.LittleEndian.PutUint64(i64LtlB, uint64(i64))
	binary.BigEndian.PutUint64(i64bigB, uint64(i64))
	fmt.Println(i64LtlB)
	fmt.Println(i64bigB)
	fmt.Println("-------")
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

	panic("123")
	fmt.Println("after panic")
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

func chg(r []rune) {
	r[0] = 100
}

/**
func c(args ...string) {
	fmt.Println(reflect.TypeOf(args))
}
**/
