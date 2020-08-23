package deck

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/deck"
	"github.com/rajankumar549/DrunkTeenPattiGame/src/util"
)

func New() deck.Deck {
	allCards := []card.Card{}
	for _, f := range card.AllAvailableCardFamily {
		for _, c := range card.AllAvailableCard {
			c.Family = f
			allCards = append(allCards, c)
		}
	}

	return &Repo{
		NoOfCardLeft: int64(len(allCards)),
		Cards:        allCards,
	}
}
func (d *Repo) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

func (d *Repo) FetchCards(howMany int64) ([]card.Card, error) {
	if d.NoOfCardLeft < howMany {
		return nil, fmt.Errorf("unable to fetch %d card cause %d card in deck", howMany, d.NoOfCardLeft)
	}
	c := d.Cards[0:howMany]
	d.Cards = d.Cards[howMany:]
	d.NoOfCardLeft -= howMany
	return c, nil
}

func (d *Repo) ShowOrder() {

	for _, c := range d.Cards {
		fmt.Printf("[%s %s]", c.Label, c.Family)
	}
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

func (repo *Repo) ReadTwoUpperCard() (card.Card, card.Card) {
	noOfCards := len(repo.Cards)
	if noOfCards <= 1 {
		return getTwoRandomCard()
	}
	noDiffcardExist := repo.IsTrail(repo.Cards)
	if noDiffcardExist {
		return getTwoRandomCard()
	}
	return repo.Cards[0], repo.Cards[1]
}

func getTwoRandomCard() (card.Card, card.Card) {
	rand1 := rand.Intn(12)
	rand2 := rand.Intn(12)
	return card.AllAvailableCard[rand1], card.AllAvailableCard[rand2]
}
