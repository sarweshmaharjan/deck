package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type log struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	/* 	bs := make([]byte,9999)
	   	resp.Body.Read(bs)

	   	fmt.Println(string(bs)) */

	lw := log{}
	// io.Copy(os.Stdout,resp.Body)
	io.Copy(lw, resp.Body)
}

func (log) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("\n byte written:", len(bs))
	return len(bs), nil
}
