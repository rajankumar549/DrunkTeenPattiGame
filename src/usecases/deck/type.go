package deck

import "github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"

type CardDeck struct {
	NoOfCardLeft int64
	Cards        []card.Card
}
