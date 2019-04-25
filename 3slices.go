package main

import "fmt"

func main() {
	cards := []string{newCard(), "5D", newCard()}
	cards = append(cards, "QD")
	fmt.Println("Cards : ", cards)

	for i, card := range cards { //for loops in go, discard 'i' and 'card' values after each iteration
		fmt.Println("Index : ", i, "||", "Card : ", card)
	}

	colors := []string{"Red", "Yellow", "Blue"}
	for col := range colors { //'col' is index, if we don't use 'i'
		fmt.Println("Color : ", colors[col]) //call element of 'colors' using 'col'
	}
}

func newCard() string { //implies we are returning a string
	return "QS"
}
