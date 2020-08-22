package player

import (
	"github.com/golang/mock/gomock"

	"github.com/rajankumar549/DrunkTeenPattiGame/src/constants"
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"
	mock "github.com/rajankumar549/DrunkTeenPattiGame/src/mock/deck"
	"testing"
)

func TestRepo_IsWinner(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	deckMock := mock.NewMockDeck(mockCtrl)
	player1 := &Repo{
		id: "P1",
		cards: []card.Card{
			card.CARD_A,
			card.CARD_A,
			card.CARD_A,
		},
	}

	player2 := &Repo{
		id: "P2",
		cards: []card.Card{
			card.CARD_A,
			card.CARD_A,
			card.CARD_2,
		},
	}
	want := constants.PlayerFaceOffResult.Win

	t.Run("P1 Winner", func(t *testing.T) {
		deckMock.EXPECT().IsTrail(player1.cards).Return(true)
		deckMock.EXPECT().IsTrail(player2.cards).Return(false)
		if got := player1.IsWinner(player2, deckMock); got != want {
			t.Errorf("IsWinner() = %v, want %v", got, want)
		}
	})
	//#1a
	player1.cards = []card.Card{
		card.CARD_A,
		card.CARD_2,
		card.CARD_3,
	}

	player2.cards = []card.Card{
		card.CARD_A,
		card.CARD_A,
		card.CARD_2,
	}
	want = constants.PlayerFaceOffResult.Tie
	t.Run("P1 Winner", func(t *testing.T) {
		deckMock.EXPECT().IsTrail(player1.cards).Return(true)
		deckMock.EXPECT().IsTrail(player2.cards).Return(true)
		if got := player1.IsWinner(player2, deckMock); got != want {
			t.Errorf("IsWinner() = %v, want %v", got, want)
		}
	})

	//#2
	player1.cards = []card.Card{
		card.CARD_A,
		card.CARD_2,
		card.CARD_3,
	}

	player2.cards = []card.Card{
		card.CARD_A,
		card.CARD_A,
		card.CARD_2,
	}
	want = constants.PlayerFaceOffResult.Win
	t.Run("P1 Winner", func(t *testing.T) {
		deckMock.EXPECT().IsTrail(player1.cards).Return(false)
		deckMock.EXPECT().IsTrail(player2.cards).Return(false)
		deckMock.EXPECT().IsValidSequence(player1.cards).Return(true)
		deckMock.EXPECT().IsValidSequence(player2.cards).Return(false)
		if got := player1.IsWinner(player2, deckMock); got != want {
			t.Errorf("IsWinner() = %v, want %v", got, want)
		}
	})

	//#2a
	player1.cards = []card.Card{
		card.CARD_A,
		card.CARD_2,
		card.CARD_3,
	}

	player2.cards = []card.Card{
		card.CARD_A,
		card.CARD_2,
		card.CARD_3,
	}
	want = constants.PlayerFaceOffResult.Tie
	t.Run("P1 Winner", func(t *testing.T) {
		deckMock.EXPECT().IsTrail(player1.cards).Return(false)
		deckMock.EXPECT().IsTrail(player2.cards).Return(false)
		deckMock.EXPECT().IsValidSequence(player1.cards).Return(true)
		deckMock.EXPECT().IsValidSequence(player2.cards).Return(true)
		if got := player1.IsWinner(player2, deckMock); got != want {
			t.Errorf("IsWinner() = %v, want %v", got, want)
		}
	})

	//#3
	player1.cards = []card.Card{
		card.CARD_A,
		card.CARD_Q,
		card.CARD_J,
	}

	player2.cards = []card.Card{
		card.CARD_A,
		card.CARD_A,
		card.CARD_2,
	}
	want = constants.PlayerFaceOffResult.Lose
	t.Run("P1 Winner", func(t *testing.T) {
		deckMock.EXPECT().IsTrail(player1.cards).Return(false)
		deckMock.EXPECT().IsTrail(player2.cards).Return(false)
		deckMock.EXPECT().IsValidSequence(player1.cards).Return(false)
		deckMock.EXPECT().IsValidSequence(player2.cards).Return(false)

		deckMock.EXPECT().HasPair(player1.cards).Return(false)
		deckMock.EXPECT().HasPair(player2.cards).Return(true)
		if got := player1.IsWinner(player2, deckMock); got != want {
			t.Errorf("IsWinner() = %v, want %v", got, want)
		}
	})

	//#3a
	player1.cards = []card.Card{
		card.CARD_A,
		card.CARD_A,
		card.CARD_J,
	}

	player2.cards = []card.Card{
		card.CARD_A,
		card.CARD_A,
		card.CARD_J,
	}
	want = constants.PlayerFaceOffResult.Tie
	t.Run("P1 Winner", func(t *testing.T) {
		deckMock.EXPECT().IsTrail(player1.cards).Return(false)
		deckMock.EXPECT().IsTrail(player2.cards).Return(false)
		deckMock.EXPECT().IsValidSequence(player1.cards).Return(false)
		deckMock.EXPECT().IsValidSequence(player2.cards).Return(false)

		deckMock.EXPECT().HasPair(player1.cards).Return(true)
		deckMock.EXPECT().HasPair(player2.cards).Return(true)
		if got := player1.IsWinner(player2, deckMock); got != want {
			t.Errorf("IsWinner() = %v, want %v", got, want)
		}
	})

	//#4
	player1.cards = []card.Card{
		card.CARD_A,
		card.CARD_2,
		card.CARD_5,
	}

	player2.cards = []card.Card{
		card.CARD_5,
		card.CARD_8,
		card.CARD_7,
	}
	want = constants.PlayerFaceOffResult.Win
	t.Run("P1 Winner", func(t *testing.T) {
		deckMock.EXPECT().IsTrail(player1.cards).Return(false)
		deckMock.EXPECT().IsTrail(player2.cards).Return(false)
		deckMock.EXPECT().IsValidSequence(player1.cards).Return(false)
		deckMock.EXPECT().IsValidSequence(player2.cards).Return(false)
		deckMock.EXPECT().HasPair(player1.cards).Return(false)
		deckMock.EXPECT().HasPair(player2.cards).Return(false)
		deckMock.EXPECT().GetTopCard(player1.cards).Return(card.CARD_A)
		deckMock.EXPECT().GetTopCard(player2.cards).Return(card.CARD_8)

		if got := player1.IsWinner(player2, deckMock); got != want {
			t.Errorf("IsWinner() = %v, want %v", got, want)
		}
	})

	//#4a
	player1.cards = []card.Card{
		card.CARD_A,
		card.CARD_2,
		card.CARD_5,
	}

	player2.cards = []card.Card{
		card.CARD_5,
		card.CARD_A,
		card.CARD_7,
	}
	want = constants.PlayerFaceOffResult.Tie
	t.Run("P1 Winner", func(t *testing.T) {
		deckMock.EXPECT().IsTrail(player1.cards).Return(false)
		deckMock.EXPECT().IsTrail(player2.cards).Return(false)
		deckMock.EXPECT().IsValidSequence(player1.cards).Return(false)
		deckMock.EXPECT().IsValidSequence(player2.cards).Return(false)
		deckMock.EXPECT().HasPair(player1.cards).Return(false)
		deckMock.EXPECT().HasPair(player2.cards).Return(false)
		deckMock.EXPECT().GetTopCard(player1.cards).Return(card.CARD_A)
		deckMock.EXPECT().GetTopCard(player2.cards).Return(card.CARD_A)
		if got := player1.IsWinner(player2, deckMock); got != want {
			t.Errorf("IsWinner() = %v, want %v", got, want)
		}
	})
}

func TestRepo_IsTieWinner(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	deckMock := mock.NewMockDeck(mockCtrl)
	player1 := &Repo{
		id: "P1",
		cards: []card.Card{
			card.CARD_A,
			card.CARD_3,
			card.CARD_5,
		},
	}

	player2 := &Repo{
		id: "P2",
		cards: []card.Card{
			card.CARD_5,
			card.CARD_3,
			card.CARD_8,
		},
	}
	want := constants.PlayerFaceOffResult.Win

	t.Run("P1 Winner", func(t *testing.T) {
		deckMock.EXPECT().GetTopCard(player1.cards).Return(card.CARD_A)
		deckMock.EXPECT().GetTopCard(player2.cards).Return(card.CARD_8)
		if got := player1.IsTieWinner(player2, deckMock); got != want {
			t.Errorf("IsWinner() = %v, want %v", got, want)
		}
	})

	//#2
	player1.cards = []card.Card{
		card.CARD_K,
		card.CARD_3,
		card.CARD_5,
	}

	player2.cards = []card.Card{
		card.CARD_A,
		card.CARD_3,
		card.CARD_8,
	}
	want = constants.PlayerFaceOffResult.Lose

	t.Run("P1 Winner", func(t *testing.T) {
		deckMock.EXPECT().GetTopCard(player1.cards).Return(card.CARD_K)
		deckMock.EXPECT().GetTopCard(player2.cards).Return(card.CARD_A)
		if got := player1.IsTieWinner(player2, deckMock); got != want {
			t.Errorf("IsWinner() = %v, want %v", got, want)
		}
	})

	//#3
	player1.cards = []card.Card{
		card.CARD_A,
		card.CARD_3,
		card.CARD_5,
	}

	player2.cards = []card.Card{
		card.CARD_A,
		card.CARD_3,
		card.CARD_8,
	}
	want = constants.PlayerFaceOffResult.Win
	t.Run("P1 Winner", func(t *testing.T) {
		deckMock.EXPECT().GetTopCard(player1.cards).Return(card.CARD_A)
		deckMock.EXPECT().GetTopCard(player2.cards).Return(card.CARD_A)
		deckMock.EXPECT().Shuffle()
		deckMock.EXPECT().ReadTwoUpperCard().Return(card.CARD_A, card.CARD_8)
		if got := player1.IsTieWinner(player2, deckMock); got != want {
			t.Errorf("IsWinner() = %v, want %v", got, want)
		}
	})

	//#5
	player1.cards = []card.Card{
		card.CARD_A,
		card.CARD_3,
		card.CARD_5,
	}

	player2.cards = []card.Card{
		card.CARD_A,
		card.CARD_3,
		card.CARD_8,
	}
	want = constants.PlayerFaceOffResult.Lose
	t.Run("P1 Winner", func(t *testing.T) {
		deckMock.EXPECT().GetTopCard(player1.cards).Return(card.CARD_A)
		deckMock.EXPECT().GetTopCard(player2.cards).Return(card.CARD_A)
		deckMock.EXPECT().Shuffle()
		deckMock.EXPECT().ReadTwoUpperCard().Return(card.CARD_A, card.CARD_A)
		deckMock.EXPECT().Shuffle()
		deckMock.EXPECT().ReadTwoUpperCard().Return(card.CARD_2, card.CARD_2)
		deckMock.EXPECT().Shuffle()
		deckMock.EXPECT().ReadTwoUpperCard().Return(card.CARD_K, card.CARD_K)
		deckMock.EXPECT().Shuffle()
		deckMock.EXPECT().ReadTwoUpperCard().Return(card.CARD_J, card.CARD_K)
		if got := player1.IsTieWinner(player2, deckMock); got != want {
			t.Errorf("IsWinner() = %v, want %v", got, want)
		}
	})
}
