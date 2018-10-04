package main

import (
	"os"
	"testing"
)

const CDeckSize = 52

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != CDeckSize {
		t.Errorf("Expected deck length of %d, but got %v", CDeckSize, len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	if d[CDeckSize-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[CDeckSize-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	d := newDeck()
	d.saveToFile("_deckTesting")
	loadedDeck := newDeckFromFile("_deckTesting")
	if len(loadedDeck) != CDeckSize {
		t.Errorf("Expected deck length of %d, but got %v", CDeckSize, len(loadedDeck))
	}
	if loadedDeck.toString() != d.toString() {
		t.Errorf("Differents decks!")
	}
	os.Remove("_deckTesting")
}
