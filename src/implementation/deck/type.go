package deck

import "github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"

type Repo struct {
	NoOfCardLeft int64
	Cards        []card.Card
}
