package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

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
}
