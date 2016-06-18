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

// NewCzechDurakGame is constructor of CzechDurakGame
func NewCzechDurakGame(player1, player2 *Player) *CzechDurakGame {
	game := new(CzechDurakGame)
	game.deck = NewRussianDeck()
	game.deck.shuffle()
	game.player1 = player1
	game.player2 = player2
	return game
}
