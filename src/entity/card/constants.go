package card

type CardFamily string

const Heart CardFamily = "♥"
const Tiels CardFamily = "♦"
const Clover CardFamily = "♠"
const Spades CardFamily = "♣"

var CARD_A = Card{
	Label:     "A",
	Seq:       1,
	FaceValue: 13,
}
var CARD_2 = Card{
	Label:     "2",
	Seq:       2,
	FaceValue: 1,
}
var CARD_3 = Card{
	Label:     "3",
	Seq:       3,
	FaceValue: 2,
}
var CARD_4 = Card{
	Label:     "4",
	Seq:       4,
	FaceValue: 3,
}
var CARD_5 = Card{
	Label:     "5",
	Seq:       5,
	FaceValue: 4,
}
var CARD_6 = Card{
	Label:     "6",
	Seq:       6,
	FaceValue: 5,
}
var CARD_7 = Card{
	Label:     "7",
	Seq:       7,
	FaceValue: 6,
}
var CARD_8 = Card{
	Label:     "8",
	Seq:       8,
	FaceValue: 7,
}
var CARD_9 = Card{
	Label:     "9",
	Seq:       9,
	FaceValue: 8,
}
var CARD_10 = Card{
	Label:     "10",
	Seq:       10,
	FaceValue: 9,
}
var CARD_J = Card{
	Label:     "J",
	Seq:       11,
	FaceValue: 10,
}
var CARD_Q = Card{
	Label:     "Q",
	Seq:       12,
	FaceValue: 11,
}
var CARD_K = Card{
	Label:     "K",
	Seq:       13,
	FaceValue: 12,
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
