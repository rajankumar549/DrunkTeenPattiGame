package card

import (
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"
	"github.com/rajankumar549/DrunkTeenPattiGame/src/util"
)

func New() card.CardHelper {
	return &Repo{}
}

func (repo *Repo) IsTrail(cards []card.Card) bool {
	noOfCards := len(cards)
	if noOfCards <= 1 {
		return false
	}

	for i := 1; i < noOfCards; i++ {
		if cards[i-1].FaceValue != cards[i].FaceValue {
			return false
		}
	}

	return true
}

func (repo *Repo) IsValidSequence(cards []card.Card) bool {
	noOfCards := len(cards)
	if noOfCards <= 1 {
		return false
	}

	copiedCards := cards
	util.SortCards(copiedCards, 1)
	for i := 1; i < noOfCards; i++ {
		if copiedCards[i-1].Seq+1 != copiedCards[i].Seq {
			return false
		}
	}
	return true
}

func (repo *Repo) HasPair(cards []card.Card) bool {
	noOfCards := len(cards)
	if noOfCards <= 1 {
		return false
	}
	copiedCards := cards
	util.SortCards(copiedCards, 1)
	for i := 1; i < noOfCards; i++ {
		//TODO :: Add Check for card Family too
		if copiedCards[i-1].FaceValue == copiedCards[i].FaceValue {
			return true
		}
	}
	return false
}

func (repo *Repo) GetTopCard(cards []card.Card) card.Card {
	noOfCards := len(cards)
	if noOfCards <= 0 {
		return card.Card{}
	} else if noOfCards == 1 {
		return cards[0]
	}
	copiedCards := cards
	util.SortCards(copiedCards, -1)
	return copiedCards[len(copiedCards)-1]
}
