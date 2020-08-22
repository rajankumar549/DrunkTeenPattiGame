package deck

import "github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"

type Deck interface {
	Shuffle()
	FetchCards(int64) ([]card.Card, error)
	ShowOrder()
	IsTrail(cards []card.Card) bool
	IsValidSequence(cards []card.Card) bool
	HasPair(cards []card.Card) bool
	GetTopCard(cards []card.Card) card.Card
	ReadTwoUpperCard() (card.Card, card.Card)
}
