package deck

import "github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"

type Deck interface {
	Shuffle()
	FetchCards(int64) ([]card.Card, error)
	ShowOrder()
}
