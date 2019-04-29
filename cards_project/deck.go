package main

import "fmt"

type deck []string //defines the type 'deck' as a slice of strings

// This is a receiver function for a deck type. It prints out all elements in a deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println("Index : ", i, "||", "Card : ", card)
	}
}
