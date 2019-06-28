package main

func main() {
	cards := newDeck()
	givenCards, remainingCards := deal(cards, 5)

	givenCards.print()
	remainingCards.print()
}
