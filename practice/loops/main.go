package main

import "fmt"

func main() {
	/*Initialize of Slice*/
	cards := []string{newCard(), "Ace of Diamonds"}

	/*To add new element. Remember 'append' return new slice and not overwrite the existing.*/
	cards = append(cards, "Six of Spades")

	/*To loop slice*/
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
