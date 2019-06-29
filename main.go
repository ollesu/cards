package main

import "fmt"

func main() {
	cards := newDeck()
	// givenCards, remainingCards := deal(cards, 5)

	fmt.Println(cards.toString())
}
