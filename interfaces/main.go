package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

type russianBot struct{ bot }

// just because those types implement [getGreeting() string] function,
// then they are considered bots
func (englishBot) getGreeting() string {
	return "hello!"
}

func (spanishBot) getGreeting() string {
	return "hola!"
}

func printGreeting(b bot) {
	fmt.Print(b.getGreeting())
}

func main() {
	eb := englishBot{}
	printGreeting(eb)

	sb := spanishBot{}
	printGreeting(sb)

	// this will break because the type does not implement getGreeting()
	// but notice I could tell the compiler russianBot type is bot interface.
	// I think this is better
	rb := russianBot{}
	printGreeting(rb)
}
