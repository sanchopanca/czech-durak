package main

// Player is an interface of players of a card game
type Player interface {
	takeCard(card Card)
	placeCard(topCard Card) *Card // TODO: It should be possible to place up to 4 cards
	handIsEmpty() bool
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

func (game *CzechDurakGame) play() {
	openCards := make([]Card, 36)
	openCards = append(openCards, game.deck.dealCard())
	players := [2]Player{*game.player1, *game.player2}
	move := 0
	for {
		player := players[move%2]
		card := player.placeCard(openCards[len(openCards)-1])
		if card == nil {
			player.takeCard(game.deck.dealCard()) // TODO: check if deck empty
		} else {
			// TODO: check is card right
			openCards = append(openCards, *card)
		}
		if player.handIsEmpty() {
			break
			// TODO: do something about winner
		}
		move++
	}
}
