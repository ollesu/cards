package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 28 {
		t.Errorf("expected deck length of 16 but got %v", len(d))
	}
}
