package main

import "fmt"

/*
* Yes, it will run correctly because the package name is same. So, printState() is refer to another package : state.go
 */
func main() {
	printState()
	/* you can call fn by just refering the name. */
	card := newCard()

	fmt.Println(card)
}
