package main

import (
	"os"
	"testing"
)

func TestDeckShouldHave52Cards(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Deck should contain 52 cards, but it contains %v", len(d))
	}
}

func TestDeckShouldStartWithAceOfSpades(t *testing.T) {
	d := newDeck()
	if d[0] != "Ace of Spades" {
		t.Errorf("Deck should start with Ace of Spades, but it starts with %v", d[0])
	}
}

func TestDeckShouldEndWithKingOfClubs(t *testing.T) {
	d := newDeck()
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Deck should end with King of Clubs, but it ends with %s", d[len(d)-1])
	}
}
func TestCreateDeck(t *testing.T) {
	os.Remove("_deckTesting")
	d := newDeck()
	d.saveToFile("_deckTesting")

	loadedDeck := loadFromFile("_deckTesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Deck should contain 52 cards, but it contains %v", len(d))
	}

	os.Remove("_deckTesting")
}
