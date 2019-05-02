package main

import "fmt"

type deck []string //defines the type 'deck' as a slice of strings

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// This is a receiver function for a deck type. It prints out all elements in a deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println("Index : ", i, "||", "Card : ", card)
	}
}

//this function returns 2 values of type 'deck'
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
