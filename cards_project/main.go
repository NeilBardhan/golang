package main

func main() {
	cards := deck{newCard(), "5D", newCard()} //Creates an object of type deck
	cards = append(cards, "QD")
	cards.print() //Calls the print method from deck
}

func newCard() string {
	return "QS"
}
