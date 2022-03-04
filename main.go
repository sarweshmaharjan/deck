/**
 * Package - It defines which project it belongs to
 */
package main

import "fmt"

/**
 * When you define package main, it needs main function to run
 * @function main
 */ 
func main(){ 
	
	// Calling function newDeck from deck.go file
	cards := newDeck()

	// Calling function deal from dock.go file with two returns that are of type deck.
	hand, remainingCards := deal(cards,5)

	// Printing val of hand and remainingCards variable to console
	hand.print()
	remainingCards.print()

	//Converting deck to string.
	all := cards.toString()
	fmt.Printf("all: %v\n", all);

	//Writing out in file. 
	cards.saveToFile("my-card-file.txt")

	//Reading from file
	newCards := newDeckFromFile("my-card-file.txt")
	newCards.print() 

	//Shuffling card
	cards.shuffle();
	cards.print();
}