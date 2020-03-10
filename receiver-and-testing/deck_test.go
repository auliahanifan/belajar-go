package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Check length of deck
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// Test the first card of array
	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card of Ace of Clubs, but got %v", d[0])
	}

	// Test the last card of array
	if d[len(d)-1] != "Four of Diamonds" {
		t.Errorf("Expected first card of Ace of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	// initialize with _ in file name, means it's temporary file
	// so please don't get mess with it
	os.Remove("_decktesting")

}
