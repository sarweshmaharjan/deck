package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	/** Channel creating */
	c := make(chan string)

	for _,link := range links{
		/** Creates new go routine if "go" keyword is used in front of it. */
		go checkLink(link,c)
	}

	/** It will loop for infinitely for{} */
	for l := range c{
	/** channel needs more print to wait for each channel request was send */
			go func(link string){
				time.Sleep(5*time.Second)
				checkLink(link,c)	
			}(l)
	}
}

func checkLink(link string,c chan string){
	_,err := http.Get(link)
	if err != nil{
		fmt.Println(link," might be down...")
		c <- link

		return
	}
	fmt.Println(link, " is up...")
	c <- link

}