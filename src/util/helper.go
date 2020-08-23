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
		if p.GetName() == winner.GetName() {
			fmt.Printf("\nName : %s (Winner) \n", p.GetName())
		} else {
			fmt.Printf("\nName : %s\n", p.GetName())
		}
		printCards(p.GetCards())
		fmt.Println("\n-------------------------------------------------")
	}
	fmt.Println("\n====================================================")
}

func printCards(inpCards []card.Card) {
	t1 := ".--------."
	t2 := "| %s.--.  |"
	t3 := "|  | %s | |"
	t4 := "|  | %s | |"
	t5 := "|  '--'%s |"
	t6 := "`--------`"
	cards := make([]string, 6)
	for _, c := range inpCards {
		cards[0] = cards[0] + "\t" + t1
		cards[1] = cards[1] + "\t" + fmt.Sprintf(t2, c.Label)
		cards[2] = cards[2] + "\t" + fmt.Sprintf(t3, c.Family)
		cards[3] = cards[3] + "\t" + fmt.Sprintf(t4, c.Family)
		cards[4] = cards[4] + "\t" + fmt.Sprintf(t5, c.Label)
		cards[5] = cards[5] + "\t" + t6
	}
	for _, s := range cards {
		fmt.Println(s)
	}
}
