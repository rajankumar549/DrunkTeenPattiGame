package player

import (
	"fmt"

	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/deck"
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/player"
)

func New(gameNo, playerNo int64, name string) player.Player {
	//do scanf here to get Actual Details
	return &Repo{
		id:   fmt.Sprintf("G%d-P%d", gameNo, playerNo),
		name: name,
	}
}

func (p *Repo) IsWinner(player2 player.Player, deck deck.Deck) player.Result {

	//case 1
	meHasValidTrail := deck.IsTrail(p.cards)
	heHasValidTrail := deck.IsTrail(player2.GetCards())
	if meHasValidTrail && heHasValidTrail {
		return player.FaceOfResult.Tie
	} else if meHasValidTrail {
		return player.FaceOfResult.Win
	} else if heHasValidTrail {
		return player.FaceOfResult.Lose
	}

	//case 2
	meHasValidSeq := deck.IsValidSequence(p.cards)
	heHasValidSeq := deck.IsValidSequence(player2.GetCards())
	if meHasValidSeq && heHasValidSeq {
		return player.FaceOfResult.Tie
	} else if meHasValidSeq {
		return player.FaceOfResult.Win
	} else if heHasValidSeq {
		return player.FaceOfResult.Lose
	}

	//case 3
	meHasValidPair := deck.HasPair(p.cards)
	heHasValidPair := deck.HasPair(player2.GetCards())
	if meHasValidPair && heHasValidPair {
		return player.FaceOfResult.Tie
	} else if meHasValidPair {
		return player.FaceOfResult.Win
	} else if heHasValidPair {
		return player.FaceOfResult.Lose
	}

	//case 4
	myTopCard := deck.GetTopCard(p.cards)
	hisTopCard := deck.GetTopCard(player2.GetCards())

	//TODO Add Card Family Check
	if myTopCard.FaceValue == hisTopCard.FaceValue {
		return player.FaceOfResult.Tie
	} else if myTopCard.FaceValue > hisTopCard.FaceValue {
		return player.FaceOfResult.Win
	} else if myTopCard.FaceValue < hisTopCard.FaceValue {
		return player.FaceOfResult.Lose
	}

	return player.FaceOfResult.Tie
}

func (p *Repo) IsTieWinner(player2 player.Player, deck deck.Deck) player.Result {

	myTopCard := deck.GetTopCard(p.cards)
	hisTopCard := deck.GetTopCard(player2.GetCards())

	if myTopCard.FaceValue > hisTopCard.FaceValue {
		return player.FaceOfResult.Win
	} else if myTopCard.FaceValue < hisTopCard.FaceValue {
		return player.FaceOfResult.Lose
	}

	return p.descideByDeckCard(deck)
}

func (p *Repo) PickCards(c []card.Card) {
	p.cards = c
}

func (p *Repo) GetCards() []card.Card {
	return p.cards
}

func (p *Repo) GetName() string {
	return fmt.Sprintf(`%s (Player ID : %s)`, p.name, p.id)
}

func (p *Repo) descideByDeckCard(deck deck.Deck) player.Result {
	deck.Shuffle()
	myTopCard, hisTopCard := deck.ReadTwoUpperCard()
	//TODO Add Card Family Check
	if myTopCard.FaceValue > hisTopCard.FaceValue {
		return player.FaceOfResult.Win
	} else if myTopCard.FaceValue < hisTopCard.FaceValue {
		return player.FaceOfResult.Lose
	}
	return p.descideByDeckCard(deck)
}
