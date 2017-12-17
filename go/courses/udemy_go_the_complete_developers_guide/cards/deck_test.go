package main

import (
	"testing"
	"os"
)

// when testing a method create a function that starts with the word 'test' and then the functions name
// uppercase function name...
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		// formatted string:  a string where you pass in some extra value and it is automatically injected at %v
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but go %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}


// func TestSaveToDeckandNewDeckFromFile() {
func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, go %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

