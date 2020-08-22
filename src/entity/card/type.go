package card

type Card struct {
	DeckID    int
	Label     string
	Seq       int
	FaceValue int
	Family    CardFamily
	Free      bool
}
