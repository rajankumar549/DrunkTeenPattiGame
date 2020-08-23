package util

import (
	"fmt"
	"testing"

	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"
)

func Test_getCardStr(t *testing.T) {

	t1 := ".--------."
	t2 := "| %s.--.  |"
	t3 := "|  (\\/)  |"
	t4 := "|  :\\/:  |"
	t5 := "|  '--'%s |"
	t6 := "`--------`"
	cards := make([]string, 6)
	t.Run("testing", func(t *testing.T) {
		for i := 1; i <= 6; i++ {
			cards[0] = cards[0] + "\t" + t1
			cards[1] = cards[1] + "\t" + fmt.Sprintf(t2, "A")
			cards[2] = cards[2] + "\t" + t3
			cards[3] = cards[3] + "\t" + t4
			cards[4] = cards[4] + "\t" + fmt.Sprintf(t5, "A")
			cards[5] = cards[5] + "\t" + t6
		}
		for _, s := range cards {
			fmt.Println(s)
		}
	})
}

func Test_printCards(t *testing.T) {
	type args struct {
		inpCards []card.Card
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				[]card.Card{
					card.CARD_A,
					card.CARD_K,
					card.CARD_Q,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printCards(tt.args.inpCards)
		})
	}
}
