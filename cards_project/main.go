package main

func main() {
	//cards := []string{newCard(), "5D", newCard()}
	cards := deck{newCard(), "5D", newCard()}
	cards = append(cards, "QD")
	// fmt.Println("Cards : ", cards)

	// for i, card := range cards { //for loops in go, discard 'i' and 'card' values after each iteration
	// 	fmt.Println("Index : ", i, "||", "Card : ", card)
	// }
	cards.print()
}

func newCard() string { //implies we are returning a string
	return "QS"
}
