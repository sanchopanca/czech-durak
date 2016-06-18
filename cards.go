package main

import (
	"math/rand"
)

// Card is one single card from a deck
type Card struct {
	value int
	suit  string
}

// RussianDeck is crad deck with 36 cards in it
// hearts, diamonds, clubs, and spades from 6 up to ace
type RussianDeck struct {
	cards []Card
}

// Hand is set of cards. All cards must be unique
type Hand struct {
	set *CardSet
}

// NewRussianDeck is constructor of RussianDeck
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

// NewHand is constructor of Hand
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
