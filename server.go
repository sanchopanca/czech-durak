package main

import (
	"fmt"
)

func main() {
	deck := NewRussianDeck()
	for _, card := range deck.cards {
		fmt.Println(card.value, card.suit)
	}
}
