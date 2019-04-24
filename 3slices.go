package main

import "fmt"

func main() {
	cards := []string{newCard(), "5D", newCard()} //a slice can only contain elements of 1 pre-defined type
	cards = append(cards, "QD") //a slice can be updated (elements added and removed), arrays are fixed length
	fmt.Println(cards) //like py lists, a slice can be printed all at once

  //can also print them itemwise
	for i, card := range cards { //for loops in go, discard 'i' and 'card' values after each iteration
		fmt.Println(i, card)
	}

	colors := []string{"Red", "Yellow", "Blue"}
	for col := range colors { //'col' is index, if we don't use 'i'
		fmt.Println(colors[col]) //call element of 'colors' using 'col'
	}
}

func newCard() string { //implies we are returning a string
	return "QS"
}
