package main

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Exepcted lenght of 20 but got %v, "len(d))
	}

	if d[0]!= "Ace of Spades"{
		t.Errorf("Expected first card Ace of spades but got %v", d[0])
	}

}


func TestSaveToDeckandNewDeckFromFile(t *testing.T){

	
	os.Remove("_decktesting")

	deck:= newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDEckFromFile("_decktesting")



	if len(loadedDeck)!=16 {
		t.Errorf("Expected 16 card in deck  but got %v", len(loadedDeck)")
	}
	os.Remove("_decktesting")

 

}
