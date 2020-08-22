package player

import "github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"

type Player interface {
	IsWinner(Player) bool
	CollectCards([]card.Card)
}
