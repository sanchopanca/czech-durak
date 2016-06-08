package main

type Player interface {
	takeCard(card Card)
	placeCard(topCard Card) Card
}
