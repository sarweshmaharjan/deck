/**
 * Let Golang know that this go file belongs to main package.
 * By doing so, we don't need to import this in main.go and directly use all the function present here.
 */
package main

/**
 * fmt - In-build package used for various helpers, such as : printing to console.
 * strings - In-build package for using helper related to strings.
 * ioutil - It gives us READ and WRITE functionality
 * os - it gives platform independent functionality.
 */
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

/**
 * Converts deck to string
 * @function toString
 * @param {deck} d
 * @return {string}
 */
func (d deck) toString() string { 
	return strings.Join([]string(d),",")
}

/**
 * Save data into file. 
 * @function saveToFile
 * @param {string} filename
 * @param {deck} d
 * @return {error}
 */
func (d deck) saveToFile(fileName string) error{ 
	return ioutil.WriteFile(fileName, []byte(d.toString()),0666)	
}

/**
 * Read from file and give {deck}. 
 * @function newDeckFromFile
 * @param {string} filename
 * @return {deck}
 */
func newDeckFromFile(fileName string) deck{	
	//err, if no err will give 'nil' else it will return error
	byteSize,err := ioutil.ReadFile(fileName)
	
	//Checking if an error was return or not.
	if err != nil{ 
		fmt.Println("Error: ",err); 
		os.Exit(1)
	}
	return deck(strings.Split(string(byteSize),","))
}

/**
 * Shuffle the deck. 
 * @function shuffle
 * @param {deck} d
 */
func (d deck) shuffle(){
	// Creating new randomize source for Rand
	source := rand.NewSource(time.Now().UnixNano())

	// Assigning source to Rand to create new Rand
	r := rand.New(source)

	for i := range d{ 
		newPosition := r.Intn(len(d) -1)
		
		// Shortcut swapping of element. 
		d[i],d[newPosition] = d[newPosition],d[i]
	}
}