package deck

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/deck"
)

func New() deck.Deck {
	allCards := []card.Card{}
	for _, f := range []card.CardFamily{card.Clover, card.Heart, card.Spades, card.Tiels} {
		for _, c := range card.AllAvailableCard {
			c.Family = f
			allCards = append(allCards, c)
		}

	}
	//fmt.Printf("All Avaliable Card are %+v", c)
	return &CardDeck{
		NoOfCardLeft: int64(len(allCards)),
		Cards:        allCards,
	}
}
func (d *CardDeck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

func (d *CardDeck) FetchCards(howMany int64) ([]card.Card, error) {
	if d.NoOfCardLeft < howMany {
		return nil, fmt.Errorf("unable to fetch %d card cause %d card in deck", howMany, d.NoOfCardLeft)
	}
	c := d.Cards[0:howMany]
	d.Cards = d.Cards[howMany:]
	d.NoOfCardLeft -= howMany
	return c, nil
}

func (d *CardDeck) ShowOrder() {
	fmt.Print("\n-------------- Deck -----------\n")
	for _, c := range d.Cards {
		fmt.Printf("[%s %s]", c.Label, c.Family)
	}
}
