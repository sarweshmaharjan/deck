/** 
 * Let Golang know that this go file belongs to main package.
 * By doing so, we don't need to import this in main.go and directly use all the function present here.
*/
package main

/** 
 * In-build package used for various helpers, such as : printing to console. 
*/
import "fmt"

/**
* Create new deck type
* @type {[]string} deck 
*/ 
type deck []string

/**
* Receiver function - Those that use type deck can use this function. 
* @function print
*/
func (d deck) print()  {
	//for-loop the deck to give out single value.
	for i,card := range d{
		fmt.Println(i,card)
	}
}

/**
 * It is used for generating cards 
 * @function newDeck 
 * @return deck 
*/
func newDeck() deck{
	cards := deck{}
	cardSuits := []string{"Spades","Diamonds","Hearts", "Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}

	// _ -> this lets Golang know that, that variable is not being used. 
	for _, suit := range cardSuits{
		for _,value := range cardValues{ 
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

/**
 * It is used to Slice range of deck. 
 * @function deal 
 * @params {deck} d
 * @params {int} handSize
 * @return deck, deck 
*/
func deal(d deck, handSize int) (deck,deck)  {
	// :handSize -> start from beginning till index-1 and handSize: -> start from handSize till the end of slice
	return d[:handSize], d[handSize:]
}