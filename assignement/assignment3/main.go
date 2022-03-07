package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main(){ 
	read := os.Args[1]
	byteSize,err := ioutil.ReadFile(string(read))
	if err != nil{ 
		fmt.Println("Error: ",err); 
		os.Exit(1)
	}
	fmt.Println(string(byteSize))
}