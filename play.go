package main

import (
	"fmt"

	"github.com/rajankumar549/DrunkTeenPattiGame/src/game"
)

func main() {
	gameCounter := int64(0)
	doYouwantTocontinue := int64(1)
	a := `___  ___  __  ___         __        __                __           __         ___  ___  ___          __       ___ ___   
 |  |__  /__'  |     \ / /  \ |  | |__)    |    |  | /  ' |__/    |__) \ /     |  |__  |__  |\ |    |__)  /\   |   |  | 
 |  |___ .__/  |      |  \__/ \__/ |  \    |___ \__/ \__, |  \    |__)  |      |  |___ |___ | \|    |    /~~\  |   |  | 
                                                                                                                        `
	fmt.Println(a)
	for doYouwantTocontinue != 0 {
		fmt.Println(`      ___          __              ___     __  ___       __  ___  ___  __  
|\ | |__  |  |    / _'  /\   |\/| |__     /__'  |   /\  |__)  |  |__  |  \ 
| \| |___ |/\|    \__> /~~\  |  | |___    .__/  |  /~~\ |  \  |  |___ |__/ 
                                                                           `)1
		fmt.Println("__________________________________________________________________________")
		gameCounter++
		g := game.New()
		g.Start(gameCounter)
		fmt.Println("Do you want to continue this this Game, for NO enter 0")
		fmt.Scanf("%d", &doYouwantTocontinue)
	}
}
