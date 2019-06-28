package main

import "fmt"

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

// this is a custom method (receiver) for type 'deck'
// any variable of the type deck in our app now gets access to print function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// d is a working instance of / reference to the deck variable
// in this example, slice cards is a d variable
