package player

import (
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/player"
)

func GetNewPlayerInfo() player.Player {
	//do scanf here to get Actual Details
	return &PlayerRepo{
		ID:   "G1-P1",
		Name: "Bob",
	}
}

func (p *PlayerRepo) IsWinner(player2 player.Player) bool {
	return true
}

func (p *PlayerRepo) CollectCards(c []card.Card) {
	p.Cards = c
}
