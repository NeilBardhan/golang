package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	cards := newDeck()
	fmt.Println("+------------ALL CARDS------------+")
	cards.print() //Calls the print method from deck
	hand, remainingDeck := deal(cards, 5)
	//A hand of size 5 and the remaining crds are returned. Both of type 'deck'
	fmt.Println("+------------CHOSEN HAND------------+")
	hand.print()
	fmt.Println("+------------REMAINING DECK------------+")
	remainingDeck.print()
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Time elapsed : ", elapsed)
}
