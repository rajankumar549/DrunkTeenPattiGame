package main

import (
	"fmt"

	"github.com/rajankumar549/DrunkTeenPattiGame/src/game"
)

func main() {
	gameCounter := int64(0)
	doYouwantTocontinue := int64(1)
	for doYouwantTocontinue != 0 {
		gameCounter++
		g := game.New()
		g.Start(gameCounter)
		fmt.Println("Do you want to continue this this Game, for NO enter 0")
		fmt.Scanf("%d", &doYouwantTocontinue)
	}
}
