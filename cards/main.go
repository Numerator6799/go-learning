package main

import (
	"fmt"
	"os"
)

func main() {
	myDeck := newDeck()
	fmt.Println(len(myDeck))
	myDeck.print()
	myDeck.shuffle()
	myDeck.print()
	hand, myDeck := myDeck.deal(5)
	fmt.Println(hand)
	// myDeck.print()
	fmt.Println(len(myDeck))
	myDeck.saveToFile("mydeck.txt")
	myDeck2 := loadFromFile("mydeck.txt")
	if myDeck2 == nil {
		fmt.Println("Could not read deck")
		fmt.Println("Quitting...")
		os.Exit(999)
	}

	myDeck2.print()
	fmt.Println(len(myDeck2))
}
