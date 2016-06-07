package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	deck := NewRussianDeck()
	// deck.shuffle()
	for deck.cardsLeft() > 0 {
		card := deck.dealCard()
		fmt.Println(card.value, card.suit)
	}
}
