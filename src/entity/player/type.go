package player

import "github.com/rajankumar549/DrunkTeenPattiGame/entity/card"

type Player struct {
	ID    string
	Name  string
	Cards [3]card.Card
}
