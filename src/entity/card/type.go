package card

type CardFamily string

type Card struct {
	DeckID    int
	Label     string
	Seq       int
	FaceValue int
	Family    CardFamily
	Free      bool
}

const Heart CardFamily = "♥"
const Tiels CardFamily = "♦"
const Clover CardFamily = "♠"
const Spades CardFamily = "♣"

var CARD_A = Card{
	Label:     "A",
	Seq:       1,
	FaceValue: 13,
	Family:    Heart,
}

var CARD_2 = Card{
	Label:     "2",
	Seq:       2,
	FaceValue: 1,
	Family:    Heart,
}
var CARD_3 = Card{
	Label:     "3",
	Seq:       3,
	FaceValue: 2,
	Family:    Heart,
}
var CARD_4 = Card{
	Label:     "4",
	Seq:       4,
	FaceValue: 3,
	Family:    Heart,
}
var CARD_5 = Card{
	Label:     "5",
	Seq:       5,
	FaceValue: 4,
	Family:    Heart,
}
var CARD_6 = Card{
	Label:     "6",
	Seq:       6,
	FaceValue: 5,
	Family:    Heart,
}
var CARD_7 = Card{
	Label:     "7",
	Seq:       7,
	FaceValue: 6,
	Family:    Heart,
}
var CARD_8 = Card{
	Label:     "8",
	Seq:       8,
	FaceValue: 7,
	Family:    Heart,
}
var CARD_9 = Card{
	Label:     "9",
	Seq:       9,
	FaceValue: 8,
	Family:    Heart,
}
var CARD_10 = Card{
	Label:     "10",
	Seq:       10,
	FaceValue: 9,
	Family:    Heart,
}
var CARD_J = Card{
	Label:     "J",
	Seq:       11,
	FaceValue: 10,
	Family:    Heart,
}
var CARD_Q = Card{
	Label:     "Q",
	Seq:       12,
	FaceValue: 11,
	Family:    Heart,
}
var CARD_K = Card{
	Label:     "K",
	Seq:       13,
	FaceValue: 12,
	Family:    Heart,
}

var AllAvailableCard = []Card{
	CARD_A,
	CARD_2,
	CARD_3,
	CARD_4,
	CARD_5,
	CARD_6,
	CARD_7,
	CARD_8,
	CARD_9,
	CARD_10,
	CARD_J,
	CARD_Q,
	CARD_K,
}

var AllAvailableCardFamily = []CardFamily{
	Heart,
	Clover,
	Spades,
	Tiels,
}
