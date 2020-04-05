package main

import "fmt"

// Create a new deck type which is a slice of strings

type deck []string

func newDeck() deck {
	var cards deck
	cardSuits := []string{"Spades", "Herts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() { // d deck is receiver i.e. any variable of type deck can access print method
	for i, card := range d {
		fmt.Println((i + 1), card)
	}
}
