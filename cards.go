package main

import (
	"math/rand"
)

type Card struct {
	value int
	suit  string
}

type RussianDeck struct {
	cards []Card
}

type Hand struct {
	set *CardSet
}

func NewRussianDeck() *RussianDeck {
	suits := []string{"hearts", "diamonds", "clubs", "spades"}

	deck := new(RussianDeck)
	deck.cards = make([]Card, 0)
	for _, suit := range suits {
		for i := 6; i <= 14; i++ {
			deck.cards = append(deck.cards, Card{i, suit})
		}
	}
	return deck
}

func NewHand() *Hand {
	hand := new(Hand)
	hand.set = NewCardSet()
	return hand
}

func (deck *RussianDeck) shuffle() {
	for i := range deck.cards {
		j := rand.Intn(i + 1)
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	}
}

func (deck *RussianDeck) dealCard() Card {
	card := deck.cards[len(deck.cards)-1]
	deck.cards = deck.cards[:len(deck.cards)-1]
	return card
}

func (deck *RussianDeck) cardsLeft() int {
	return len(deck.cards)
}

func (hand *Hand) addCard(card Card) {
	hand.set.add(card)
}

func (hand *Hand) removeCard(card Card) {
	hand.set.remove(card)
}
