package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// i need to create a new type => deck
// it is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Spades", "Diamond", "Hearts", "Carre"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuits {
		for _, value := range cardValues {
			cards = append(cards, value+"of"+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
