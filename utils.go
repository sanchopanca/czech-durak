package main

type CardSet struct {
	set map[Card]bool
}

func NewCardSet() *CardSet {
	cardSet := new(CardSet)
	cardSet.set = make(map[Card]bool)
	return cardSet
}

func (cardSet *CardSet) has(card Card) bool {
	return cardSet.set[card]
}

func (cardSet *CardSet) add(card Card) {
	cardSet.set[card] = true
}

func (cardSet *CardSet) remove(card Card) {
	if cardSet.set[card] {
		delete(cardSet.set, card)
	}
}
