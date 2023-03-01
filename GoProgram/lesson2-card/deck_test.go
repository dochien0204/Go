package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	length := len(d)

	if length != 20 {
		t.Errorf("Expected length of deck 16 but got %v", length)
	}

}
