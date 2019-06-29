package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// create a new type called 'deck'
// (which is / has a behaviour of a slice of strings)
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// this is a custom method (receiver function) for type 'deck'
// any variable of the type deck in our app now gets access to print function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// d is a working instance of / reference to the deck variable
// in this example, slice cards is a d variable

// (deck, deck) refers to several return value types from the function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// type we want - type we have (d)
	string := []string(d)
	return strings.Join(string, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
