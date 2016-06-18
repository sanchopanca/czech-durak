package main

// CardSet is set of unique cards
type CardSet struct {
	set map[Card]bool
}

// NewCardSet is constructor of CardSet
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
