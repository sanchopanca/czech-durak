package main

// Player is an interface of players of a card game
type Player interface {
	takeCard(card Card)
	placeCard(topCard Card) Card
}

// CzechDurakGame is rules of the game
type CzechDurakGame struct {
	deck    *RussianDeck
	player1 *Player
	player2 *Player
}
