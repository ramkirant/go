package main

func main() {
	//cards := []string{newCard(), newCard()}
	//cards := newDeck()

	//hand, _ := deal(cards, 5)

	//hand.saveToFile("mycards.txt")

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
