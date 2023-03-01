package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	length := len(d)

	if length != 16 {
		t.Errorf("Expected length of deck 16 but got %v", length)
	}
}

func TestSaveToFileAndDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := deckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected length of card 16 from file but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
