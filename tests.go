package main

import "fmt"

//* it's possible to declare variables outside a function...
var a string = "hi"
var b int
var c int = 20

//* ...but the below variable declaration is not valid outside a function
// d := 30

func main() {
	var deckSize int
	deckSize = 52
	b = 50
	fmt.Println(a)
	fmt.Println(deckSize)
	fmt.Println(b)
	fmt.Println(c)

	//* those variable declararios are the same (var name type = value) (name := value)
	//var card string = "Ace of Spades"
	card := "Ace of Spades"

	//* ':=' operator is just for variable declaration, for assignment use the regular '=' operator
	card = "Ace of Hearts"

	fmt.Println(card)

	fmt.Println(getValue())

	//* printState function is declared in the file state.go (but need to run 'go run tests.go state.go' to include all files)
	//printState()

	//* getting the return variables of function Println, I can use the discard ('_') operator
	n, _ := fmt.Println("test")
	fmt.Println(n)

	//* declaring a slice
	cards := []string{
		"Ace of Spades",
		"Five of Hearts",
	}
	cards = append(cards, "Six of Spades")

	//* there are 2 ways of iteracting over a slice
	for i := 0; i < len(cards); i++ {
		fmt.Println(i, cards[i])
	}
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

//* return type of a function is right after the function's name and parameters declaration
func getValue() string {
	return "something"
}
