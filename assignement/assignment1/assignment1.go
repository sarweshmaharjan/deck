package main

import "fmt"

func main() {
	val := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, item := range val {
		if item%2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}
	}
}
