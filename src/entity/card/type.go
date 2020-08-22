package card

type Card struct {
	DeckID    int
	Label     string
	Seq       int
	FaceValue int
	Family    CardFamily
	Free      bool
}

type CardHelper interface {
	IsTrail(cards []Card) bool
	IsValidSequence(cards []Card) bool
	HasPair(cards []Card) bool
	GetTopCard(cards []Card) Card
}
