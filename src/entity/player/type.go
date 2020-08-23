package player

import (
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/deck"
)

type Result int64

type Player interface {
	IsWinner(Player, deck.Deck) Result
	IsTieWinner(Player, deck.Deck) Result
	PickCards([]card.Card)
	GetCards() []card.Card
	GetName() string
}

var FaceOfResult = struct {
	Win  Result
	Lose Result
	Tie  Result
}{
	Win:  1,
	Lose: -1,
	Tie:  0,
}
