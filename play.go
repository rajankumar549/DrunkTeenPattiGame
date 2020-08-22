package DrunkTeenPattiGame

import (
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/player"
	"github.com/rajankumar549/DrunkTeenPattiGame/src/usecases/deck"
)

func main() {
	//n := 2
	//Get All Player Name
	p1 := player.Player{
		ID:   "G1-P1",
		Name: "Bob",
	}
	p2 := player.Player{
		ID:   "G1-P2",
		Name: "Jhon",
	}

	deck1 := deck.New()
	deck1.Shuffle()
	deck1.Shuffle()

	// Assign Card to Each Player
	p1.Cards, _ = deck1.FetchCards(3)
	p2.Cards, _ = deck1.FetchCards(3)
	// CalCaulateWinner
}
