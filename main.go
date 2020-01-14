package main

func main() {

	// cards := newDeck()
	// cards.saveToFile("MyCards")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
