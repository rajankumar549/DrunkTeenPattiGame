package deck

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/deck"
)

func New() deck.Deck {
	c := []card.Card{}
	for _, f := range []card.CardFamily{card.Clover, card.Heart, card.Spades, card.Tiels} {
		c = append(c, card.Card{
			Label:     "A",
			Seq:       1,
			FaceValue: 13,
			Family:    f,
		}, card.Card{
			Label:     "2",
			Seq:       2,
			FaceValue: 1,
			Family:    f,
		},
			card.Card{
				Label:     "3",
				Seq:       3,
				FaceValue: 2,
				Family:    f,
			},
			card.Card{
				Label:     "4",
				Seq:       4,
				FaceValue: 3,
				Family:    f,
			},
			card.Card{
				Label:     "5",
				Seq:       5,
				FaceValue: 4,
				Family:    f,
			},
			card.Card{
				Label:     "6",
				Seq:       6,
				FaceValue: 5,
				Family:    f,
			},
			card.Card{
				Label:     "7",
				Seq:       7,
				FaceValue: 6,
				Family:    f,
			},
			card.Card{
				Label:     "8",
				Seq:       8,
				FaceValue: 7,
				Family:    f,
			},
			card.Card{
				Label:     "9",
				Seq:       9,
				FaceValue: 8,
				Family:    f,
			},
			card.Card{
				Label:     "10",
				Seq:       10,
				FaceValue: 9,
				Family:    f,
			},
			card.Card{
				Label:     "J",
				Seq:       11,
				FaceValue: 10,
				Family:    f,
			},
			card.Card{
				Label:     "Q",
				Seq:       12,
				FaceValue: 11,
			},
			card.Card{
				Label:     "K",
				Seq:       13,
				FaceValue: 12,
				Family:    f,
			})
	}
	//fmt.Printf("All Avaliable Card are %+v", c)
	return &CardDeck{
		NoOfCardLeft: int64(len(c)),
		Cards:        c,
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
