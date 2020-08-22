package player

import "github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"

type PlayerRepo struct {
	ID    string
	Name  string
	Cards []card.Card
}
