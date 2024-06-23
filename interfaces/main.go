package main

import "fmt"

type englishBot struct {
}

type spanishBot struct {
}

func (englishBot) getGreeting() string {
	return "hello!"
}

func (spanishBot) getGreeting() string {
	return "hola!"
}

func printGreeting(eb englishBot) {
	fmt.Print(eb.getGreeting())
}

func printGreeting2(sb spanishBot) {
	fmt.Print(sb.getGreeting())
}

func main() {
	eb := englishBot{}
	printGreeting(eb)

	sb := spanishBot{}
	printGreeting2(sb)
}
