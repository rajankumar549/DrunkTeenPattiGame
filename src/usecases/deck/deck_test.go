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
