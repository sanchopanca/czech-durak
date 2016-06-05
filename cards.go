package main

type Card struct {
	value int
	suit  string
}

type RussianDeck struct {
	cards []Card
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
