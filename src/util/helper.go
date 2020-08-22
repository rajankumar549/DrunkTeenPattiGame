package util

import (
	"fmt"
	"sort"

	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/player"
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

func DeclareResult(winner player.Player, players []player.Player) {
	fmt.Println("\n====================================================")
	for _, p := range players {
		cardStr := getcardStr(p.GetCards())
		if p.GetName() == winner.GetName() {
			fmt.Printf("\nName : %s  Cards : %s (Winner)", p.GetName(), cardStr)
		} else {
			fmt.Printf("\nName : %s  Cards : %s", p.GetName(), cardStr)
		}

	}
	fmt.Println("\n====================================================")
}

func ShowPlayersCard(player []player.Player) {
	fmt.Println("\n====================Other Player Details================================\n")
	for _, p := range player {
		cardStr := getcardStr(p.GetCards())
		fmt.Printf("\nName : %s  Cards : %s\n", p.GetName(), cardStr)
	}
	fmt.Println("\n====================End================================\n")
}

func getcardStr(cards []card.Card) string {
	cardDetails := ""
	for _, c := range cards {
		cardDetails = fmt.Sprintf("%s [%s %s]", cardDetails, c.Label, c.Family)
	}
	return cardDetails
}
