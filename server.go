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
	for i := 0; i < 36; i++ {
		card := deck.getCard()
		fmt.Println(card.value, card.suit)
	}
}
