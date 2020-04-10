package main

func main() {
	cards := newDeck()

	_, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// fmt.Println(cards.toString())
	remainingCards.saveToFile("my_cards")
	cards = newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}

// go run main.go deck.go
