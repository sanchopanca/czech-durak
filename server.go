package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	deck := NewRussianDeck()
	hand := NewHand()
	// deck.shuffle()
	for deck.cardsLeft() > 0 {
		card := deck.dealCard()
		fmt.Println(hand.set.has(card))
		hand.addCard(card)
		fmt.Println(hand.set.has(card))
		hand.removeCard(card)
		fmt.Println(hand.set.has(card))
		fmt.Println(card.value, card.suit)
		return
	}
}
