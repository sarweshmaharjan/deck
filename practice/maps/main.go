package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#745654",
	}

	addresses := make(map[string]string)
	addresses["home"] = "Paknajol,kathmandu,nepal"
	addresses["work"] = "somewhere"

	fmt.Println(addresses)

	delete(addresses, "work")

	printMap(addresses)
	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for index, address := range c {
		fmt.Printf("The address is %s in position %s \n", address, index)
	}
}
