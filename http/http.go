package main

import (
	"fmt"
	"io"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!")
	req.ParseForm()
	if len(req.Form) > 0 {
		for k, v := range req.Form {
			fmt.Printf("%s, %s\n", k, v[0])
		}
	}
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println(err)
	}
}
