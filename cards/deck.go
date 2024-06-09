package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv" // I was using this to convert integer to string (Itoa function)
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	//var d deck // alternative to declare deck
	d := deck{}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	specialValues := []string{"Jack", "Queen", "King"}
	for _, suit := range suits {
		d = append(d, createCard("Ace", suit))
		for i := 2; i <= 10; i++ {
			d = append(d, createCard(strconv.Itoa(i), suit))
		}
		for _, sv := range specialValues {
			d = append(d, createCard(sv, suit))
		}
	}
	return d
}

func createCard(cardValue string, cardSuit string) string {
	return cardValue + " of " + cardSuit
}

// * function with a receiver
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) deal(handSize int) ([]string, deck) {
	h := d[:handSize] // or 0:handsize
	d = d[handSize:]  //* this doesn't change the list parameter, we have to return the deck
	return h, d
}

func (d deck) saveToFile(path string) {
	//*I was using this code to create the content but the toString from the course method is better
	// var content string = ""
	// for _, card := range d {
	// 	content += card + ","
	// }

	os.WriteFile(path, []byte(d.toString()), 0644) // this is unix permissions
}

func (d deck) toString() string {
	//* course said to convert deck to []string but there's no need as deck is already a slice os strings
	return strings.Join(d, ",")
}

func loadFromFile(path string) deck {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("An error ocurred", err)
		return nil
	}
	content := string(data)
	var d deck = strings.Split(content, ",")
	return d
}

func (d deck) shuffle() {
	// Seed the random number generator
	// rand.Seed(time.Now().UnixNano()) // no need to call Seed anymore, do the below instead
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Fisher-Yates shuffle algorithm
	for i := len(d) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}
