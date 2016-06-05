package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	deck := NewRussianDeck()
	deck.shuffle()
	for _, card := range deck.cards {
		fmt.Println(card.value, card.suit)
	}
}
