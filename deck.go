package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

type deck []string

func newDeck() deck {
	var cards deck
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
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

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return newDeck()
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() deck {
	var temp string

	for i := range d {
		random := rand.Intn(len(d) - 1) // rand.Intn is a pseudo random generator, same everytime(Check documentation)
		temp = d[i]
		d[i] = d[random]
		d[random] = temp
	}

	return d
}
