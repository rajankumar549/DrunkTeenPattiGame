package player

import (
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/deck"
)

type Player interface {
	IsWinner(Player, deck.Deck) int64
	IsTieWinner(Player, deck.Deck) int64

	PickCards([]card.Card)
	GetCards() []card.Card
	GetName() string
}
