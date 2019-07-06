package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 28 {
		t.Errorf("expected deck length of 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("expected first card to be Ace of Spades but got %v", d[0])
	}

	// len(d) - 1 refers to the last element of the slice
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("expected last card to be King of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 28 {
		t.Errorf("Expected 28 cards in the deck, got %d", len(loadedDeck))
	}
	os.Remove("_decktesting")

}
