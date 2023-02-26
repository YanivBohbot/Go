package main

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Exepcted lenght of 20 but got , "len(d))
	}

}
