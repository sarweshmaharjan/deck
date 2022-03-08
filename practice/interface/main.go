package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBots struct{}

func main() {

	eb := englishBot{}
	sb := spanishBots{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBots) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
