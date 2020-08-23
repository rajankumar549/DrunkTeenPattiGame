package game

import (
	"fmt"

	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/deck"
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/player"
	deckRepo "github.com/rajankumar549/DrunkTeenPattiGame/src/implementation/deck"
	player2 "github.com/rajankumar549/DrunkTeenPattiGame/src/implementation/player"
	"github.com/rajankumar549/DrunkTeenPattiGame/src/util"
)

const players_Hand_Card_Size = 3

type Game struct {
	Deck    deck.Deck
	Players []player.Player
}

func (gm *Game) Start(gameNo int64) {
	n := gm.getNoOfPlayers()
	gm.AddNPlayerToGame(gameNo, n)

	gm.Deck.Shuffle()
	gm.Deck.Shuffle()

	for i := int64(0); i < n; i++ {
		thisPlayerCards, err := gm.Deck.FetchCards(players_Hand_Card_Size)
		if err != nil {
			fmt.Printf("Something went wrong :: Unable to fetch %d cards for %s", players_Hand_Card_Size, gm.Players[i].GetName())
			return
		}
		gm.Players[i].PickCards(thisPlayerCards)
	}
	winner := gm.Players[0]
	if n == 1 {
		util.DeclareResult(winner, gm.Players)
		return
	}

	for i := int64(1); i < n; i++ {

		result := winner.IsWinner(gm.Players[i], gm.Deck)
		if result == player.FaceOfResult.Win {
			continue
		}

		if result == player.FaceOfResult.Lose {
			winner = gm.Players[i]
			continue
		}

		tieResult := winner.IsTieWinner(gm.Players[i], gm.Deck)
		if tieResult == player.FaceOfResult.Lose {
			winner = gm.Players[i]
		}
	}
	util.DeclareResult(winner, gm.Players)
	return
}

func New() *Game {
	return &Game{
		Deck: deckRepo.New(),
	}
}

func (gm *Game) AddNPlayerToGame(gameNo, n int64) {
	for i := int64(0); i < n; i++ {
		var name = ""
		fmt.Printf("\nEnter Player %d Name : ", i+1)
		fmt.Scanf("%s", &name)
		p := player2.New(gameNo, i, name)
		gm.Players = append(gm.Players, p)
		fmt.Printf("%s Susseccfully Joined", p.GetName())
	}
}

func (gm *Game) getNoOfPlayers() int64 {
	notValid := true
	n := 0
	for notValid {
		fmt.Print("\nEnter How many Player going to play this Game. : ")
		fmt.Scanf("%d", &n)
		if (players_Hand_Card_Size*n) > 52 || n <= 0 {
			fmt.Printf("Invalid No Of Players , Right now we can allow only %d players at a time \n", 52/players_Hand_Card_Size)
			notValid = true

		} else {
			notValid = false
		}
	}
	return int64(n)
}
