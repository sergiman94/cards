package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 16 {
		t.Errorf("Expected length of 16 but got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades as first item but got %v", deck[0])
	}

	if deck[len(deck)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs as last item but got %v", deck[len(deck)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected to have length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
