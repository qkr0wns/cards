package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length %v, but got %v", 52, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card %v, but got %v", "Ace of Spades", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card %v, but got %v", "King of Clubs", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length %v, but got %v", 52, len(loadedDeck))
	}
	os.Remove("_decktesting")
}
