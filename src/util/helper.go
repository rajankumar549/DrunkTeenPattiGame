package util

import (
	"sort"

	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"
)

func SortCards(cards []card.Card, opt int) {
	switch opt {
	case -1:
		sort.SliceStable(cards, func(i, j int) bool {
			return cards[i].FaceValue < cards[j].FaceValue
		})
	default:
		sort.SliceStable(cards, func(i, j int) bool {
			return cards[i].Seq < cards[j].Seq
		})
	}
}
