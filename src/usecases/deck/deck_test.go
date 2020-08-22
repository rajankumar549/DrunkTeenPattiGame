package deck

import (
	"reflect"
	"testing"

	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/card"
	"github.com/rajankumar549/DrunkTeenPattiGame/src/entity/deck"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want deck.Deck
	}{
		{
			name: "show All Deck",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			if got == nil {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardDeck_Shuffle(t *testing.T) {
	deck1 := New()
	deck2 := New()
	t.Run("Shuffling 1", func(t *testing.T) {
		deck1.Shuffle()
		deck2.Shuffle()
		if reflect.DeepEqual(deck1, deck2) {
			deck1.ShowOrder()
			deck2.ShowOrder()
			t.Errorf("Both Deck Have Same Card Even After Suffling")
		}
	})
	t.Run("Shuffling 1", func(t *testing.T) {
		deck1.Shuffle()
		deck2.Shuffle()
		if reflect.DeepEqual(deck1, deck2) {
			deck1.ShowOrder()
			deck2.ShowOrder()
			t.Errorf("Both Deck Have Same Card Even After Suffling")
		}
	})
}

func TestCardDeck_FetchCards(t *testing.T) {
	deck1 := New()
	t.Run("Fetch # cards", func(t *testing.T) {
		got, err := deck1.FetchCards(1)
		if err != nil {
			t.Errorf("unable to fetch card %v", err)
		}
		want := []card.Card{
			{
				Label:     "A",
				Family:    card.Clover,
				Seq:       1,
				FaceValue: 13,
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v cards but got %v cards", want, got)
		}
	})
	t.Run("Fetching 49", func(t *testing.T) {
		got, err := deck1.FetchCards(51)
		if err != nil {
			t.Errorf("unable to fetch card %v", err)
		}

		if len(got) < 49 {
			t.Errorf("Inavlid Cards got %+v", got)
		}
	})

	t.Run("Fetching on Empty Deck", func(t *testing.T) {
		_, err := deck1.FetchCards(1)
		if err == nil {
			deck1.ShowOrder()
			t.Error("Should throw erron on fetching empty deck")
		}
	})
}

func TestRepo_IsTrail(t *testing.T) {
	type args struct {
		cards []card.Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Invalid Trail",
			args: args{
				cards: []card.Card{},
			},
		},
		{
			name: "Invalid Trail",
			args: args{
				cards: []card.Card{
					card.CARD_2,
				},
			},
		},
		{
			name: "Invalid Trail",
			args: args{
				cards: []card.Card{
					card.CARD_2,
					card.CARD_4,
					card.CARD_2,
				},
			},
		},
		{
			name: "Valid Trail",
			args: args{
				cards: []card.Card{
					card.CARD_2,
					card.CARD_2,
					card.CARD_2,
				},
			},
			want: true,
		},
		{
			name: "Invalid Trail",
			args: args{
				cards: []card.Card{
					card.CARD_3,
					card.CARD_4,
					card.CARD_2,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repo{}
			if got := repo.IsTrail(tt.args.cards); got != tt.want {
				t.Errorf("IsTrail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_IsValidSequence(t *testing.T) {
	type args struct {
		cards []card.Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Invalid Seq",
			args: args{
				cards: []card.Card{},
			},
		},
		{
			name: "Invalid Seq",
			args: args{
				cards: []card.Card{
					card.CARD_2,
				},
			},
		},
		{
			name: "Invalid Seq",
			args: args{
				cards: []card.Card{
					card.CARD_2,
					card.CARD_3,
					card.CARD_5,
				},
			},
		},
		{
			name: "Valid Seq",
			args: args{
				cards: []card.Card{
					card.CARD_A,
					card.CARD_2,
					card.CARD_3,
				},
			},
			want: true,
		},
		{
			name: "Valid Seq",
			args: args{
				cards: []card.Card{
					card.CARD_3,
					card.CARD_4,
					card.CARD_2,
				},
			},
			want: true,
		},
		{
			name: "Valid Seq",
			args: args{
				cards: []card.Card{
					card.CARD_J,
					card.CARD_K,
					card.CARD_Q,
				},
			},
			want: true,
		},
		{
			name: "Valid Seq",
			args: args{
				cards: []card.Card{
					card.CARD_10,
					card.CARD_J,
					card.CARD_Q,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repo{}
			if got := repo.IsValidSequence(tt.args.cards); got != tt.want {
				t.Errorf("IsValidSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_HasPair(t *testing.T) {
	type args struct {
		cards []card.Card
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Has No Pair",
			args: args{
				cards: []card.Card{},
			},
		},
		{
			name: "Has No Pair",
			args: args{
				cards: []card.Card{
					card.CARD_2,
				},
			},
		},
		{
			name: "Has No Pair",
			args: args{
				cards: []card.Card{
					card.CARD_2,
					card.CARD_3,
					card.CARD_5,
				},
			},
		},
		{
			name: "Has Pair",
			args: args{
				cards: []card.Card{
					card.CARD_A,
					card.CARD_2,
					card.CARD_A,
				},
			},
			want: true,
		},
		{
			name: "Has Pair",
			args: args{
				cards: []card.Card{
					card.CARD_3,
					card.CARD_3,
					card.CARD_2,
				},
			},
			want: true,
		},
		{
			name: "Valid Seq",
			args: args{
				cards: []card.Card{
					card.CARD_J,
					card.CARD_K,
					card.CARD_K,
				},
			},
			want: true,
		},
		{
			name: "Valid Seq",
			args: args{
				cards: []card.Card{
					card.CARD_Q,
					card.CARD_J,
					card.CARD_Q,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repo{}
			if got := repo.HasPair(tt.args.cards); got != tt.want {
				t.Errorf("HasPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_GetTopCard(t *testing.T) {
	type args struct {
		cards []card.Card
	}
	tests := []struct {
		name string
		args args
		want card.Card
	}{
		{
			name: "Invalid Seq",
			args: args{
				cards: []card.Card{},
			},
		},
		{
			name: "Invalid Seq",
			args: args{
				cards: []card.Card{
					card.CARD_2,
				},
			},
			want: card.CARD_2,
		},
		{
			name: "Card 5",
			args: args{
				cards: []card.Card{
					card.CARD_2,
					card.CARD_3,
					card.CARD_5,
				},
			},
			want: card.CARD_5,
		},
		{
			name: "Card A",
			args: args{
				cards: []card.Card{
					card.CARD_A,
					card.CARD_2,
					card.CARD_3,
				},
			},
			want: card.CARD_A,
		},
		{
			name: "Card Q",
			args: args{
				cards: []card.Card{
					card.CARD_3,
					card.CARD_Q,
					card.CARD_2,
				},
			},
			want: card.CARD_Q,
		},
		{
			name: "Card",
			args: args{
				cards: []card.Card{
					card.CARD_K,
					card.CARD_K,
					card.CARD_Q,
				},
			},
			want: card.CARD_K,
		},
		{
			name: "Valid Seq",
			args: args{
				cards: []card.Card{
					card.CARD_10,
					card.CARD_Q,
					card.CARD_Q,
				},
			},
			want: card.CARD_Q,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repo{}
			if got := repo.GetTopCard(tt.args.cards); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTopCard() = %v, want %v", got, tt.want)
			}
		})
	}
}
